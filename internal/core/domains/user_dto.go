package domains

import (
	"time"

	"github.com/google/uuid"
)

//------ User ------ //

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email,lowercase" example:"example@example.local"`
	Password string `json:"password" validate:"required" example:"12345678910"`
}
type RegisterDTO struct {
	FirstName string  `json:"first_name" validate:"required" example:"Resul"`
	LastName  string  `json:"last_name" validate:"required" example:"Çelik"`
	Lang      string  `json:"lang" example:"tr"` //🤘
	Email     string  `json:"email" validate:"required,email,lowercase" example:"example@example.local"`
	Password  string  `json:"password" validate:"required,min=10" example:"12345678910"`
	Gender    string  `json:"gender" validate:"required,oneof=male female not other" example:"male"`
	Appeal    *string `json:"appeal" validate:"omitempty" example:"I am a Attack Helicopter 🚁"`
}

type SessionDTO struct {
	ID             *uuid.UUID `db:"id"`
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

type LoginResponseDTO struct {
	ID            *uuid.UUID `db:"id" json:"-"`
	UpdatedAt     *time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt     *time.Time `db:"created_at" json:"created_at"`
	Password      string     `db:"password" json:"password,omitempty"`
	FirstName     string     `db:"first_name" json:"firstname" validate:"required,alphaunicode" example:"Resul"`
	LastName      string     `db:"last_name" json:"lastname" validate:"required,alphaunicode" example:"Çelik"`
	Lang          string     `db:"lang" json:"lang" example:"tr"` //🤘
	Email         string     `db:"email" json:"email" validate:"omitempty,email" example:"resul@mon.time"`
	EmailVerified bool       `db:"email_verified" json:"email_verified"`
	MasterAdmin   bool       `db:"master_admin" json:"master_admin"`
	Banned        bool       `db:"banned" json:"-"`
	Gender        string     `db:"gender" json:"gender" validate:"required,oneof=male female not other" example:"male"`
	Appeal        *string    `db:"appeal" json:"appeal" validate:"omitempty" example:"I am a Attack Helicopter 🚁"`
	MFAEnabled    bool       `db:"mfa_enabled" json:"mfa_enabled"`
	Disabled      bool       `db:"disabled" json:"-"`
	DisabledAt    *time.Time `db:"disabled_at" json:"-"`
	LastLogin     *time.Time `db:"last_login" json:"last_login"`
}
