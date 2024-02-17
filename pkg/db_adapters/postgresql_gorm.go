package db_adapters

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Veri tabanı bağlantısını başlatacak fonksiyon
func NewGormClient(host, port, username, password, database string) (*gorm.DB, error) {

	// Veri tabanı bağlantısı başlat
	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=" + host + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable",
		PreferSimpleProtocol: true, // JSON, JSONB, hstore, timestamptz, and interval,s
	}), &gorm.Config{})

}
