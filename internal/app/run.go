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
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/error_handler"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/middlewares"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/server"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/store"
	"github.com/un-defined-gsc/un-defined-backend/internal/repositories"
	"github.com/un-defined-gsc/un-defined-backend/pkg/db_adapters"
	"github.com/un-defined-gsc/un-defined-backend/pkg/validator_service"
)

func Run(cfg *config.Config) {
	//postgreClient
	db, err := db_adapters.NewGormClient(cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Database)
	rdb, err := db_adapters.NewRedisClient(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password)
	if err != nil {
		panic(err)
	}

	// repository initialize
	userRepo := repositories.NewUserRepositories(pool, rdb)

	// email service initialize
	emailService := email.EmailInit(cfg.Email.Address, cfg.Email.Name, cfg.Email.Host, cfg.Email.Port, cfg.Email.Username, cfg.Email.Password)
	go emailService.WriteStdoutError() //doğru bir yöntem değil

	// service initialize
	deps := deps_services.NewDepsServices(captcha_service.Init(rdb), hasher_service.NewHasherService(), emailService, otp_serivce.NewOTPService(config.GetConfig().App.Site, 30), validator_service.NewValidatorService())
	userser := user_services.NewUsersServices(userRepo, deps)
	dataser := data_services.NewDataServices(cfg.Data.Host, cfg.Data.Port)
	monitorser := monitor_services.NewMonitorServices(monitorRepo, dataser, deps)
	feedbackser := feedback_services.NewFeedbackServices(feedbackRepo, deps)

	// adapter initialize
	adapter := core.NewCoreAdapter(userser, feedbackser, deps, monitorser)

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
