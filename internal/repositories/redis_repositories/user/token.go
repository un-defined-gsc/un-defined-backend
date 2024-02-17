package user_redis_repositories

import (
	"context"

	"github.com/redis/go-redis/v9"
	user_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/user"
	user_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/user"
)

type tokensRepository struct {
	redis *redis.Client
}

func NewTokensRepository(redis *redis.Client) user_ports.ITokensRepository {
	return &tokensRepository{
		redis: redis,
	}
}

func (p *tokensRepository) Create(ctx context.Context, token *user_domain.Token) (err error) {
	//Transaction redis
	p.redis.Watch(ctx, func(tx *redis.Tx) error {
		pipe := tx.TxPipeline()
		pipe.Set(ctx, token.Token, token, token.ExpiresAt.Sub(token.CreatedAt))
		_, err = pipe.Exec(ctx)
		return err
	}, token.Token)
	return
}

func (p *tokensRepository) GetByToken(ctx context.Context, token string) (user *user_domain.Token, err error) {
	user = new(user_domain.Token)
	p.redis.Get(ctx, token).Scan(user)
	return
}

func (p *tokensRepository) DeleteByToken(ctx context.Context, token string) (err error) {
	err = p.redis.Del(ctx, token).Err()
	return
}
