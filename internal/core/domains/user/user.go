package user_domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            *uuid.UUID `db:"id" json:"-"`
	UpdatedAt     *time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt     *time.Time `db:"created_at" json:"created_at"`
	Password      string     `db:"password" json:"password,omitempty"`
	FirstName     string     `db:"first_name" json:"firstname" validate:"required,alphaunicode" example:"Resul"`
	LastName      string     `db:"last_name" json:"lastname" validate:"required,alphaunicode" example:"Çelik"`
	Lang          string     `db:"lang" json:"lang" validate:"required,oneof=en tr" example:"tr"` //🤘
	Email         string     `db:"email" json:"email" validate:"omitempty,email" example:"resul@mon.time"`
	EmailVerified bool       `db:"email_verified" json:"email_verified"`
	MasterAdmin   bool       `db:"master_admin" json:"-"`
	Banned        bool       `db:"banned" json:"-"`
	Gender        string     `db:"gender" json:"gender" validate:"required,oneof=male female not other" example:"male"`
	Appeal        *string    `db:"appeal" json:"appeal" validate:"omitempty" example:"I am a Attack Helicopter 🚁"`
	MFAEnabled    bool       `db:"mfa_enabled" json:"mfa_enabled"`
	Disabled      bool       `db:"disabled" json:"-"`
	DisabledAt    *time.Time `db:"disabled_at" json:"-"`
	LastLogin     *time.Time `db:"last_login" json:"last_login"`
}

func (u User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
