package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/un-defined-gsc/un-defined-backend/internal/config"
	"github.com/un-defined-gsc/un-defined-backend/internal/core"
	deps_services "github.com/un-defined-gsc/un-defined-backend/internal/core/services/deps"
	user_services "github.com/un-defined-gsc/un-defined-backend/internal/core/services/user"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/error_handler"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/middlewares"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/server"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/store"
	"github.com/un-defined-gsc/un-defined-backend/internal/repositories"
	gorm_migration "github.com/un-defined-gsc/un-defined-backend/internal/repositories/gorm_repositories"
	censor_service "github.com/un-defined-gsc/un-defined-backend/pkg/censor"
	"github.com/un-defined-gsc/un-defined-backend/pkg/db_adapters"
	hasher_service "github.com/un-defined-gsc/un-defined-backend/pkg/hasher"
	email "github.com/un-defined-gsc/un-defined-backend/pkg/mailler"
	otp_serivce "github.com/un-defined-gsc/un-defined-backend/pkg/otp"
	"github.com/un-defined-gsc/un-defined-backend/pkg/validator_service"
)

func Run(cfg *config.Config) {
	//gormclient
	db, err := db_adapters.NewGormClient(cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Database)
	if err != nil {
		panic(err)
	}
	rdb, err := db_adapters.NewRedisClient(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password)
	if err != nil {
		panic(err)
	}

	// repository initialize
	userRepo := repositories.NewUserRepositories(db, rdb)

	if err := gorm_migration.StartMigration(); err != nil {
		panic(err)
	}

	// email service initialize
	emailService := email.EmailInit(cfg.Email.Address, cfg.Email.Name, cfg.Email.Host, cfg.Email.Port, cfg.Email.Username, cfg.Email.Password)
	go emailService.WriteStdoutError() //doğru bir yöntem değil

	// service initialize
	deps := deps_services.NewDepsServices(
		hasher_service.NewHasherService(),
		emailService, otp_serivce.NewOTPService(config.GetConfig().App.Site, 30),
		validator_service.NewValidatorService(),
		censor_service.NewCensorService("./locales/badwords.txt"),
	)
	userser := user_services.NewUsersServices(userRepo, deps)

	// adapter initialize
	adapter := core.NewCoreAdapter(userser, deps)

	//handler initialize
	handlers := http.NewHandler(adapter)

	//server initialize
	fiberSrv := server.NewServer(cfg, error_handler.ErrorHandler)

	//fiber store initialize
	fiberStore := store.NewFiberStore(cfg.Redis.Host, cfg.Redis.Password, cfg.Redis.Port)

	//captcha store initialize

	go func() {
		err := fiberSrv.Run(handlers.Init(cfg, fiberStore, middlewares.InitMiddlewares(cfg, fiberStore)...))
		if err != nil {
			log.Fatalf("Error while running fiber server: %v", err.Error())
		}
	}()
	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	<-c                                             // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")      // Daha iyi yapılabilir !!
	_ = fiberSrv.Shutdown(context.Background())
	fmt.Println("Fiber was successful shutdown.")
}

// func databaseMigrate(migrationPath string, pool *pgxpool.Pool) error {
// 	err := goose.SetDialect("postgres")
// 	if err != nil {
// 		return err
// 	}
// 	db := stdlib.OpenDBFromPool(pool)
// 	if err := goose.Up(db, migrationPath); err != nil {
// 		return err
// 	}
// 	if err := db.Close(); err != nil {
// 		return err
// 	}
// 	return nil
// }
