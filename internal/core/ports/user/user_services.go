package user_ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	user_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/user"
)

type IUsersService interface {
	Login(ctx context.Context, login domains.LoginDTO) (sess *domains.SessionDTO, userdata user_domain.User, err error)
	EnableSession(ctx context.Context, session *domains.SessionDTO, token string) (err error)
	Register(ctx context.Context, user domains.RegisterDTO) (err error)
	VerifyProfile(ctx context.Context, token string) (err error)
	VerifyEmail(ctx context.Context, token string) (err error)
	MFAToggle(ctx context.Context, userID uuid.UUID) (status bool, err error)
	GetMFA(ctx context.Context, userID uuid.UUID) (mfa *user_domain.MFASetting, err error)
	GetMe(ctx context.Context, userID uuid.UUID) (user *user_domain.User, err error)
	UpdateMe(ctx context.Context, user user_domain.User) (err error)
	ChangeEmail(ctx context.Context, newEmail domains.EmailCahangeDTO) (err error)
	// ResendVerifyEmail(ctx context.Context, newEmail domains.EmailCahangeDTO) (err error) Şimdilik gerek yok çünkü rediste tutuluyor ama ileride bi düzenleme gerekecek
	ChangePassword(ctx context.Context, newPassword domains.PasswordChangeDTO) (err error)
	SendRecoveryToken(ctx context.Context, email domains.PasswordRecoverySendDTO) (err error)
	RecoverPassword(ctx context.Context, newPassword domains.PasswordRecoveryDTO) (err error)
}

type IUsersServices interface {
	UsersService() IUsersService
}
