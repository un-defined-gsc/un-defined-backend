package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	user_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/user"

	user_ps_repositories "github.com/un-defined-gsc/un-defined-backend/internal/repositories/postgresql_repositories/user"
	user_redis_repositories "github.com/un-defined-gsc/un-defined-backend/internal/repositories/redis_repositories/user"
)

type userRepositories struct {
	usersRepository    user_ports.IUsersRepository
	bannedsRepository  user_ports.IBannedsRepository
	mfasRepository     user_ports.IMFAsRepository
	tokenRepository    user_ports.ITokensRepository
	tempUserRepository user_ports.ITempUsersRepository
}

func NewUserRepositories(dbpool *pgxpool.Pool, redis *redis.Client) user_ports.IUsersRepositories {
	return &userRepositories{
		usersRepository:    user_ps_repositories.NewUsersRepository(dbpool),
		bannedsRepository:  user_ps_repositories.NewBannedsRepository(dbpool),
		mfasRepository:     user_ps_repositories.NewMFAsRepository(dbpool),
		tokenRepository:    user_redis_repositories.NewTokensRepository(redis),
		tempUserRepository: user_redis_repositories.NewTempUsersRepository(redis),
	}
}

func (r *userRepositories) UsersRepository() user_ports.IUsersRepository {
	return r.usersRepository
}

func (r *userRepositories) BannedsRepository() user_ports.IBannedsRepository {
	return r.bannedsRepository
}

func (r *userRepositories) MFAsRepository() user_ports.IMFAsRepository {
	return r.mfasRepository
}

func (r *userRepositories) TokensRepository() user_ports.ITokensRepository {
	return r.tokenRepository
}

func (r *userRepositories) TempUsersRepository() user_ports.ITempUsersRepository {
	return r.tempUserRepository
}
