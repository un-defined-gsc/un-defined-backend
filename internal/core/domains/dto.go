package domains

import (
	"time"

	"github.com/google/uuid"
	social_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/social"
)

//------ User ------ //

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email,lowercase" example:"rsuresulcelik@gmail.com"`
	Password string `json:"password" validate:"required" example:"12345678"`
}
type RegisterDTO struct {
	FirstName string `json:"first_name" validate:"required" example:"Resul"`
	LastName  string `json:"last_name" validate:"required" example:"Çelik"`
	Lang      string `json:"lang" validate:"required,oneof=en tr" example:"tr"` //🤘
	Email     string `json:"email" validate:"required,email,lowercase" example:"rsuresulcelik@resulcelik.net"`
	Password  string `json:"password" validate:"required,min=10" example:"12345678910"`
}

type SessionDTO struct {
	UUID           *uuid.UUID `db:"uuid"`
	EnabledSession bool       `db:"enabled_session"`
	Email          string     `db:"email"`
	FirstName      string     `db:"first_name"`
	LastName       string     `db:"last_name"`
	EmailVerified  bool       `db:"email_verified"`
	MasterAdmin    bool       `db:"master_admin"`
	Lang           string     `db:"lang"`
	MFAEnabled     bool       `db:"mfa_enabled"`
	Disabled       bool       `db:"disabled"`
	DisabledAt     *time.Time `db:"disabled_at"`
	LastLogin      *time.Time `db:"last_login"`
	CreatedAt      *time.Time `db:"created_at"`
	Key            *string    `db:"key"`
}

type PasswordChangeDTO struct {
	UserId             uuid.UUID `json:"-"`
	OldPassword        string    `json:"old_password" validate:"required,min=10" example:"12345678"`
	NewPassword        string    `json:"new_password" validate:"required,min=10,nefield=OldPassword" example:"12345678"`
	NewPasswordConfirm string    `json:"new_password_confirm" validate:"required,min=10,eqfield=NewPassword" example:"12345678"`
}

type EmailCahangeDTO struct {
	UserId   uuid.UUID `json:"-"`
	Email    string    `json:"email" validate:"required,email,lowercase" example:"test@example.com"`
	Password string    `json:"password" validate:"required" example:"12345678"`
}

type PasswordRecoverySendDTO struct {
	Email string `json:"email" validate:"required,email,lowercase" example:"rsuresulcelik@gmail.com"`
}

type PasswordRecoveryDTO struct {
	Token              string `json:"-"`
	NewPassword        string `json:"new_password" validate:"required,min=10" example:"12345678"`
	NewPasswordConfirm string `json:"new_password_confirm" validate:"required,min=10,eqfield=NewPassword" example:"12345678"`
}

//----- Social ------//

type CategoryDTO struct {
	ID   uuid.UUID `gorm:"primary_key" json:"id"`
	Name string    `json:"category"`
}


type  PostDTO struct {
	ID          uuid.UUID `gorm:"column:id" json:"id"`
	social_domain.Category  
