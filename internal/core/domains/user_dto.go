package domains

import (
	"time"

	"github.com/google/uuid"
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
	ID             *uuid.UUID `gorm:"column:id"`
	EnabledSession bool       `gorm:"column:enabled_session"`
	Email          string     `gorm:"column:email"`
	FirstName      string     `gorm:"column:first_name"`
	LastName       string     `gorm:"column:last_name"`
	EmailVerified  bool       `gorm:"column:email_verified"`
	MasterAdmin    bool       `gorm:"column:master_admin"`
	Lang           string     `gorm:"column:lang"`
	MFAEnabled     bool       `gorm:"column:mfa_enabled"`
	Disabled       bool       `gorm:"column:disabled"`
	DisabledAt     *time.Time `gorm:"column:disabled_at"`
	LastLogin      *time.Time `gorm:"column:last_login"`
	CreatedAt      *time.Time `gorm:"column:created_at"`
	Key            *string    `gorm:"column:key"`
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
