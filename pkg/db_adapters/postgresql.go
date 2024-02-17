package db_adapters

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Veri tabanı bağlantısını başlatacak fonksiyon
func NewPostgressClient(host, port, username, password, database string) (*pgxpool.Pool, error) {

	ctx := context.Background()

	// Veri tabanı bağlantısı başlat
	dbpool, err := pgxpool.New(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database))
	if err != nil {
		return nil, err
	}

	// Bağlantıyı test etmek için ping fonksiyonu çağırılır.
	err = dbpool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return dbpool, nil
}
