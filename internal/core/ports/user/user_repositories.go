package user_ports

import (
	"context"
	"time"

	users 
	"github.com/google/uuid"
)

type IUsersRepository interface {
	// User Table Commands //
	Create(ctx context.Context, user *users.User) (err error)
	Update(ctx context.Context, newUser *users.User) (err error)
	UpdateEmailVerifiedTrueByUUID(ctx context.Context, userUUID uuid.UUID) (err error)
	UpdateDisabledByUUID(ctx context.Context, userUUID uuid.UUID, newDisabled bool) (err error)
	UpdateMasterAdminByUUID(ctx context.Context, userUUID uuid.UUID, newMasterAdmin bool) (err error)
	UpdatePasswordByUUID(ctx context.Context, userUUID uuid.UUID, newPassword string) (err error)
	UpdateEmailByUUID(ctx context.Context, userUUID uuid.UUID, newEmail string) (err error)
	UpdateLastLoginByUUID(ctx context.Context, userUUID uuid.UUID) (err error)
	// End User Table Commands //

	// User Table Queries //
	GetByUUID(ctx context.Context, userUUID uuid.UUID) (user *users.User, err error)
	GetByEmail(ctx context.Context, email string) (user *users.User, err error)
	FindAll(ctx context.Context, searchUser users.User, limit, offset uint64) (users []users.User, count uint64, err error)
	GetCountByEmail(ctx context.Context, email string) (count uint64, err error)
	// End User Table Queries //

	// UpdateRegisterEmailUUIDByUUID(ctx context.Context, userUUID uuid.UUID, newRegisterMailUUID uuid.UUID) (err error)
	// UpdateUsernameByUUID(ctx context.Context, userUUID uuid.UUID, newUsername string) (err error)
}

// -------------------------------------------------------------------------------------------------------------//
type IBannedsRepository interface {
	// Banned Table Commands //
	Create(ctx context.Context, banned *users.Banned) (err error)
	Update(ctx context.Context, newBanned *users.Banned) (err error)
	DeleteByUUID(ctx context.Context, bannedUUID uuid.UUID) (err error)
	DeleteByUserUUID(ctx context.Context, userUUID uuid.UUID) (err error)
	// End Banned Table Commands //
	// Banned Table Queries //
	GetByUserUUID(ctx context.Context, userUUID uuid.UUID) (banned *users.Banned, err error)
	// End Banned Table Queries //
}

// -------------------------------------------------------------------------------------------------------------//
type ITokensRepository interface {
	// Recovery Redis Commands //
	Create(ctx context.Context, token *users.Token) (err error)
	DeleteByToken(ctx context.Context, token string) (err error)
	// End Recovery Redis Commands //

	// Recovery Redis Queries //
	GetByToken(ctx context.Context, token string) (tokenResult *users.Token, err error)
	// End Recovery Redis Queries //
}

type ITempUsersRepository interface {
	// Recovery Redis Commands //
	Create(ctx context.Context, key string, tempUser *users.User) (expDate time.Time, err error)
	DeleteByKey(ctx context.Context, key string) (err error)
	GetByKey(ctx context.Context, key string) (tempUser *users.User, err error)
	// End Recovery Redis Commands //
}

// -------------------------------------------------------------------------------------------------------------//
type IMFAsRepository interface {
	// MFA Ops
	Create(ctx context.Context, mfaSetting *users.MFASetting) (err error)
	Update(ctx context.Context, newMFASetting *users.MFASetting) (err error)
	DeleteByUserUUID(ctx context.Context, userUUID uuid.UUID) (err error)
	UpdateLogUUIDByUUID(ctx context.Context, mfaSettingUUID, newLogUUID uuid.UUID) (err error)
	// End MFA Ops

	// MFA Table Queries //
	GetByUserUUID(ctx context.Context, userUUID uuid.UUID) (mfaSetting *users.MFASetting, err error)
	// End MFA Table Queries //
}

// -------------------------------------------------------------------------------------------------------------//
type IUsersRepositories interface {
	UsersRepository() IUsersRepository
	BannedsRepository() IBannedsRepository
	MFAsRepository() IMFAsRepository
	TokensRepository() ITokensRepository
	TempUsersRepository() ITempUsersRepository
}

// İlk etapta devre dışı bırakıldı. tüm alanlar user modeline taşındı.

// -------------------------------------------------------------------------------------------------------------//
// type IProfilesRepository interface { // Profile modeli devre dışı bırakıldı. Bazı alanları user modeline taşındı.
// 	// Profile Table Commands //
// 	Upsert(ctx context.Context, newProfile *users.Profile) (err error)
// 	// End Profile Table Commands //

// 	// Profile Table Queries //
// 	GetByUserUUID(ctx context.Context, userUUID uuid.UUID) (profile *users.Profile, err error)
// 	// End Profile Table Queries //
// }

//-------------------------------------------------------------------------------------------------------------//
// type IAddressesRepository interface {
// 	// Address Table Commands //
// 	Create(ctx context.Context, address *users.Address) (err error)
// 	DeleteByUUID(ctx context.Context, addressUUID uuid.UUID) (err error)
// 	Update(ctx context.Context, newAddress *users.Address) (err error)
// 	// End Address Table Commands //
// 	// Address Table Queries //
// 	FindAll(ctx context.Context, searchAddress users.Address, limit, offset uint64) (addresses []users.Address, count uint64, err error)
// 	// End Address Table Queries //
// }

// -------------------------------------------------------------------------------------------------------------//

// -------------------------------------------------------------------------------------------------------------//
// type IEmailsRepository interface {
// 	// Mail Table Commands //
// 	Create(ctx context.Context, email *users.Email) (err error)
// 	DeleteByUUID(ctx context.Context, emailUUID uuid.UUID) (err error)
// 	UpdatePrimaryEmailByUserUUIDAndUUID(ctx context.Context, userUUID, emailUUID uuid.UUID) (err error)
// 	UpdatePrimaryAndVerifedLogUUIDByUUID(ctx context.Context, emailUUID, newVerifiedLogUUID uuid.UUID) (err error)
// 	// End Mail Table Commands //

// 	// Emails Table Queries //
// 	FindAll(ctx context.Context, searchEmail users.Email, limit, offset uint64) (emails []users.Email, count uint64, err error)
// 	// End Emails Table Queries //

// }

// -------------------------------------------------------------------------------------------------------------//
// type IPhonesRepository interface {
// 	// Phone Table Commands //
// 	Create(ctx context.Context, phone *users.Phone) (err error)
// 	DeleteByUUID(ctx context.Context, phoneUUID uuid.UUID) (err error)
// 	UpdatePrimaryPhoneByUserUUIDAndUUID(ctx context.Context, userUUID, phoneUUID uuid.UUID) (err error)
// 	UpdatePrimaryAndVerifedLogUUIDByUUID(ctx context.Context, phoneUUID, newVerifiedLogUUID uuid.UUID) (err error)
// 	// End Phone Table Commands //

// 	// Phones Table Queries //
// 	FindAll(ctx context.Context, searchPhone users.Phone, limit, offset uint64) (phones []users.Phone, count uint64, err error)
// 	// End Phones Table Queries //
// }

// -------------------------------------------------------------------------------------------------------------//
// type IAuthsRepository interface {

// 	// JOIN Queries //
// 	GetByUsername(ctx context.Context, username string) (userAuth *users.LoginAuthResponseDTO, err error)
// 	/* JOIN:
// 	- users
// 	- mfa
// 	- phones
// 	- emails
// 	*/
// }

// -------------------------------------------------------------------------------------------------------------//
// type IUsersAdminRepository interface {
// 	FindAll(ctx context.Context, searchUser users.User, limit, offset uint64) (users []users.UserAdminDTO, count uint64, err error)
// 	/* JOIN:
// 	-user
// 	-profile
// 	-mail (primary)
// 	-phone (primary)
// 	*/
// }
