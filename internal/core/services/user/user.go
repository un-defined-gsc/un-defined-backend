package user_services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/un-defined-gsc/un-defined-backend/internal/config"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	user_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/user"
	service_errors "github.com/un-defined-gsc/un-defined-backend/internal/core/errors"
	deps_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/deps"
	user_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/user"
	randomstr "github.com/un-defined-gsc/un-defined-backend/pkg/random_str"
)

type userService struct {
	userRepositories user_ports.IUsersRepositories
	deps             deps_ports.IDepsServices
}

func newUsersService(
	userRepositories user_ports.IUsersRepositories,
	deps deps_ports.IDepsServices,
) user_ports.IUsersService {
	return &userService{
		userRepositories: userRepositories,
		deps:             deps,
	}
}

func (s *userService) Login(ctx context.Context, login domains.LoginDTO) (sess *domains.SessionDTO, err error) {
	if err = s.deps.ValidatorService().ValidateStruct(login); err != nil {
		return
	}
	user, err := s.userRepositories.UsersRepository().GetByEmail(ctx, login.Email)
	if err != nil {
		if err == service_errors.ErrDataNotFound {
			err = service_errors.ErrInvalidUsernameOrPassword
		}
		return
	}
	if user == nil {
		err = service_errors.ErrInvalidUsernameOrPassword
		return
	}
	stat, err := s.deps.HasherService().CompareHashAndPassword(user.Password, login.Password)
	if err != nil {
		return
	}
	if !stat {
		err = service_errors.ErrInvalidUsernameOrPassword
		return
	}
	// if !user.EmailVerified {
	// 	err = service_errors.ErrUserEmailNotVerified
	// 	return
	// }
	if user.Disabled {
		err = service_errors.ErrUserIsDisabled
		return
	}
	if user.Banned {
		err = service_errors.ErrUserIsBanned
		return
	}
	sess = &domains.SessionDTO{
		ID:             user.ID,
		EnabledSession: true,
		Email:          user.Email,
		EmailVerified:  user.EmailVerified,
		MasterAdmin:    user.MasterAdmin,
		Lang:           user.Lang,
		MFAEnabled:     user.MFAEnabled,
		Disabled:       user.Disabled,
		DisabledAt:     user.DisabledAt,
		LastLogin:      user.LastLogin,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		CreatedAt:      user.CreatedAt,
	}
	if user.MFAEnabled {
		mfaSet, err := s.userRepositories.MFAsRepository().GetByUserID(ctx, *user.ID)
		if err != nil {
			return nil, err
		}
		sess.Key = mfaSet.Key
		sess.EnabledSession = false
	}
	if err = s.userRepositories.UsersRepository().UpdateLastLoginByID(ctx, *user.ID); err != nil {
		return
	}
	nowTime := time.Now().UTC()
	sess.LastLogin = &nowTime
	return
}

func (s *userService) Register(ctx context.Context, register domains.RegisterDTO) (err error) {
	if err = s.deps.ValidatorService().ValidateStruct(register); err != nil {
		return
	}
	count, err := s.userRepositories.UsersRepository().GetCountByEmail(ctx, register.Email)
	if err != nil {
		return
	}
	if count > 0 {
		// Burada error dönmüyoruz
		return
	}
	hash, err := s.deps.HasherService().HashPassword(register.Password)
	if err != nil {
		return
	}
	user := &user_domain.User{
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Lang:      register.Lang,
		Email:     register.Email,
		Password:  hash,
		Gender:    register.Gender,
		Appeal:    register.Appeal,
	}
	key := randomstr.RandStringBytesMaskImpr(48)

	expTime, err := s.userRepositories.TempUsersRepository().Create(ctx, key, user)
	if err != nil {
		return err
	}
	link := fmt.Sprintf("http://%s/auth/verify/init/%s", config.GetConfig().App.Site, key) // TODO: confige göre ayarlanacak
	err = s.deps.MailService().SendRegisterVerifyMail(ctx, user.FirstName, user.LastName, expTime.UTC().Format("2006-01-02 15:04:05"), user.Email, link)
	return
}

func (s *userService) VerifyProfile(ctx context.Context, token string) (err error) {
	user, err := s.userRepositories.TempUsersRepository().GetByKey(ctx, token)

	if err != nil {
		if err == service_errors.ErrDataNotFound {
			return service_errors.ErrTokenInvalid
		}
		return
	}

	_, err = s.userRepositories.UsersRepository().GetByEmail(ctx, user.Email)
	if err != nil {
		if err != service_errors.ErrDataNotFound {
			return
		}
	}
	user.EmailVerified = true
	if err = s.userRepositories.UsersRepository().Create(ctx, user); err != nil {
		return
	}
	if err = s.userRepositories.TempUsersRepository().DeleteByKey(ctx, token); err != nil {
		return
	}
	return
}

func (s *userService) VerifyEmail(ctx context.Context, token string) (err error) {
	newMail, err := s.userRepositories.TokensRepository().GetByToken(ctx, token)
	if err != nil {
		if err == service_errors.ErrDataNotFound {
			return service_errors.ErrTokenInvalid
		}
		return

	}
	err = s.userRepositories.UsersRepository().UpdateEmailByID(ctx, newMail.UserID, newMail.NewEmail)
	if err != nil {
		return
	}
	// err = s.userRepositories.UsersRepository().UpdateEmailVerifiedTrueByUUID(ctx, user.UserUUID) //Zaten eklerken doğrulanmış olucak
	// if err != nil {
	// 	return
	// }
	if err = s.userRepositories.TokensRepository().DeleteByToken(ctx, token); err != nil {
		return
	}
	// eski maile bildirim gönderilecek
	err = s.deps.MailService().SendNotifyOldMail(ctx, newMail.FirstName, newMail.LastName, time.Now().UTC().Format("2006-01-02 15:04:05"), newMail.OldMail, true)
	if err != nil {
		return
	}
	//Kodlar buraya
	return
}

func (s *userService) ChangeEmail(ctx context.Context, newEmail domains.EmailCahangeDTO) (err error) { //Bunlar doğru kullanım mı bilmiyorum şimdilik böyle kalsın
	if err = s.deps.ValidatorService().ValidateStruct(newEmail); err != nil {
		return
	}
	user, err := s.userRepositories.UsersRepository().GetByID(ctx, newEmail.UserId)
	if err != nil {
		return
	}
	stat, err := s.deps.HasherService().CompareHashAndPassword(user.Password, newEmail.Password)
	if err != nil {
		return
	}
	if !stat {
		err = service_errors.ErrInvalidUsernameOrPassword
		return
	}
	key := randomstr.RandStringBytesMaskImpr(48)
	err = s.userRepositories.TokensRepository().Create(ctx, &user_domain.Token{
		UserID:    newEmail.UserId,
		NewEmail:  newEmail.Email,
		OldMail:   user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Token:     key,
		ExpiresAt: time.Now().UTC().Add(time.Minute * 15),
		CreatedAt: time.Now().UTC(),
	})
	if err != nil {
		return
	}
	link := fmt.Sprintf("http://%s/auth/verify/email/%s", config.GetConfig().App.Site, key) // TODO: confige göre ayarlanacak
	err = s.deps.MailService().SendMailChangeVerify(ctx, user.FirstName, user.LastName, time.Now().UTC().Add(time.Minute*15).Format("2006-01-02 15:04:05"), newEmail.Email, link)
	if err != nil {
		return
	}
	// Eski maile bildirim gönderilecek
	err = s.deps.MailService().SendNotifyOldMail(ctx, user.FirstName, user.LastName, time.Now().UTC().Format("2006-01-02 15:04:05"), user.Email, false)
	if err != nil {
		return
	}
	//Kodlar buraya
	return
}

// func (s *userService) ResendVerifyEmail(ctx context.Context, newEmail domains.EmailCahangeDTO) (err error) { //Bunlar doğru kullanım mı bilmiyorum şimdilik böyle kalsın
// 	key := randomstr.RandStringBytesMaskImpr(48)
// 	err = s.userRepositories.TokensRepository().Create(ctx, &user_domain.Token{
// 		UserUUID:  newEmail.UserId,
// 		Token:     key,
// 		ExpiresAt: time.Now().Add(time.Minute * 15),
// 		CreatedAt: time.Now(),
// 	})
// 	if err != nil {
// 		return
// 	}
// 	link := fmt.Sprintf("http://%s/auth/verify/email/%s", config.GetConfig().App.Site, key) // TODO: confige göre ayarlanacak
// 	err = s.deps.MailService().SendRegisterVerifyMail(ctx, newEmail.FirstName, newEmail.LastName, time.Now().Add(time.Minute*15).Format("2006-01-02 15:04:05"), newEmail.Email, link)
// 	return
// }

func (s *userService) ChangePassword(ctx context.Context, password domains.PasswordChangeDTO) (err error) {
	if err = s.deps.ValidatorService().ValidateStruct(password); err != nil {
		return
	}
	user, err := s.userRepositories.UsersRepository().GetByID(ctx, password.UserId)
	if err != nil {
		return
	}
	stat, err := s.deps.HasherService().CompareHashAndPassword(user.Password, password.OldPassword)
	if err != nil {
		return
	}
	if !stat {
		err = service_errors.ErrInvalidOldPassword
		return
	}
	hash, err := s.deps.HasherService().HashPassword(password.NewPassword)
	if err != nil {
		return
	}
	return s.userRepositories.UsersRepository().UpdatePasswordByID(ctx, password.UserId, hash)
}

func (s *userService) GetMe(ctx context.Context, userID uuid.UUID) (user *user_domain.User, err error) {
	user, err = s.userRepositories.UsersRepository().GetByID(ctx, userID)
	user.Password = ""
	return
}

func (s *userService) UpdateMe(ctx context.Context, newUser user_domain.User) (err error) {
	err = s.deps.ValidatorService().ValidateStruct(newUser)
	if err != nil {
		return
	}
	err = s.userRepositories.UsersRepository().Update(ctx, &newUser)
	if err != nil {
		return
	}
	return
}

func (s *userService) SendRecoveryToken(ctx context.Context, email domains.PasswordRecoverySendDTO) (err error) {
	if err = s.deps.ValidatorService().ValidateStruct(email); err != nil {
		return
	}
	user, err := s.userRepositories.UsersRepository().GetByEmail(ctx, email.Email)
	if err != nil {
		if err == service_errors.ErrDataNotFound {
			err = nil
		}
		return
	}
	token := randomstr.RandStringBytesMaskImpr(48)
	err = s.userRepositories.TokensRepository().Create(ctx, &user_domain.Token{
		UserID:    *user.ID,
		Token:     token,
		ExpiresAt: time.Now().UTC().Add(time.Minute * 15),
		CreatedAt: time.Now().UTC(),
	})
	if err != nil {
		return
	}
	link := fmt.Sprintf("http://%s/auth/recover/password/%s", config.GetConfig().App.Site, token) // TODO: confige göre ayarlanacak
	err = s.deps.MailService().SendPasswordRecoveryMail(ctx, user.FirstName, user.LastName, time.Now().UTC().Add(time.Minute*15).Format("2006-01-02 15:04:05"), user.Email, link)
	return
}

func (s *userService) RecoverPassword(ctx context.Context, newPassword domains.PasswordRecoveryDTO) (err error) {
	if err = s.deps.ValidatorService().ValidateStruct(newPassword); err != nil {
		return
	}
	user, err := s.userRepositories.TokensRepository().GetByToken(ctx, newPassword.Token)
	if err != nil {
		if err == service_errors.ErrDataNotFound {
			return service_errors.ErrTokenInvalid
		}
		return
	}
	hash, err := s.deps.HasherService().HashPassword(newPassword.NewPassword)
	if err != nil {
		return
	}
	err = s.userRepositories.UsersRepository().UpdatePasswordByID(ctx, user.UserID, hash)
	if err != nil {
		return
	}
	if err = s.userRepositories.TokensRepository().DeleteByToken(ctx, newPassword.Token); err != nil {
		return
	}
	return
}

func (s *userService) EnableSession(ctx context.Context, session *domains.SessionDTO, token string) (err error) {
	stat, err := s.deps.OTPService().ValidateOTP(token, *session.Key)
	if err != nil {
		return
	}
	if !stat {
		err = service_errors.ErrInvalidOTP
	}
	session.EnabledSession = true
	return
}

func (s *userService) MFAToggle(ctx context.Context, userID uuid.UUID) (status bool, err error) {
	user, err := s.userRepositories.UsersRepository().GetByID(ctx, userID)
	if err != nil {
		return
	}
	if user.MFAEnabled {
		err = s.userRepositories.MFAsRepository().DeleteByUserID(ctx, userID)
		if err != nil {
			return
		}
		status = false
	} else {
		otpURL, err := s.deps.OTPService().GenerateOTP(user.Email)
		if err != nil {
			return false, err
		}
		t := time.Now().UTC()
		err = s.userRepositories.MFAsRepository().Create(ctx, &user_domain.MFASetting{
			UserID:    *user.ID,
			Key:       &otpURL,
			CreatedAt: &t,
		})
		if err != nil {
			return false, err
		}
		status = true
	}
	return
}

func (s *userService) GetMFA(ctx context.Context, userID uuid.UUID) (mfa *user_domain.MFASetting, err error) {
	mfa, err = s.userRepositories.MFAsRepository().GetByUserID(ctx, userID)
	if err != nil {
		if err == service_errors.ErrDataNotFound {
			err = service_errors.ErrMFANotEnabled
		}
		return
	}
	mfa.KeyImage, err = s.deps.OTPService().GenerateImage(*mfa.Key)
	return
}
