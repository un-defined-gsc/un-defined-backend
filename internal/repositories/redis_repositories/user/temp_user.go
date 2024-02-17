package user_redis_repositories

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	user_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/user"
	service_errors "github.com/un-defined-gsc/un-defined-backend/internal/core/errors"
	user_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/user"
)

type TempUsersRepository struct {
	redis *redis.Client
}

func NewTempUsersRepository(redis *redis.Client) user_ports.ITempUsersRepository {
	return &TempUsersRepository{
		redis: redis,
	}
}

func (r *TempUsersRepository) Create(ctx context.Context, key string, tempUser *user_domain.User) (expTime time.Time, err error) {
	resp := r.redis.Set(ctx, key, tempUser, time.Hour*3) // 3 hours sonra temp kullanıcı silinecek bunun burada ayarlanması uygun değil admin panelden ayarlanması lazım ama şimdilik böyle
	err = resp.Err()
	expTime = time.Now().UTC().Add(time.Hour * 3)
	return
}

func (r *TempUsersRepository) DeleteByKey(ctx context.Context, key string) (err error) {
	resp := r.redis.Del(ctx, key)
	err = resp.Err()
	if err != nil {
		if err == redis.Nil {
			err = service_errors.ErrDataNotFound
		}
	}

	return
}

func (r *TempUsersRepository) GetByKey(ctx context.Context, key string) (tempUser *user_domain.User, err error) {
	resp := r.redis.Get(ctx, key)
	err = resp.Err()
	if err != nil {
		if err == redis.Nil {
			err = service_errors.ErrDataNotFound
		}
		return
	}
	tempUser = new(user_domain.User)
	err = resp.Scan(tempUser)
	return
}
