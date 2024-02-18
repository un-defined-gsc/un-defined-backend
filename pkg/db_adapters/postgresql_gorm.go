package db_adapters

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Veri tabanı bağlantısını başlatacak fonksiyon
func NewGormClient(host, port, username, password, database string) (*gorm.DB, error) {
	// Veri tabanı bağlantısı başlat
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=" + host + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable",
		PreferSimpleProtocol: true, // JSON, JSONB, hstore, timestamptz, and interval,s
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return db, err
	}
	err = db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
	if err != nil {
		return db, err
	}
	return db, nil
}
