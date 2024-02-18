package user_domain

import (
	"encoding/json"
	"time"

	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

type User struct {
	base_domain.Base
	Password      string     `gorm:"column:password" json:"password,omitempty"`
	FirstName     string     `gorm:"column:first_name" json:"firstname" validate:"required,alphaunicode" example:"Resul"`
	LastName      string     `gorm:"column:last_name" json:"lastname" validate:"required,alphaunicode" example:"Çelik"`
	Lang          string     `gorm:"column:lang" json:"lang" validate:"required,oneof=en tr" example:"tr"` //🤘
	Email         string     `gorm:"column:email,unique" json:"email" validate:"omitempty,email" example:"resul@mon.time"`
	EmailVerified bool       `gorm:"column:email_verified" json:"email_verified"`
	MasterAdmin   bool       `gorm:"column:master_admin" json:"-"`
	Banned        bool       `gorm:"column:banned" json:"-"`
	MFAEnabled    bool       `gorm:"column:mfa_enabled" json:"mfa_enabled"`
	UpdatedAt     *time.Time `gorm:"column:updated_at" json:"-"`
	Disabled      bool       `gorm:"column:disabled" json:"-"`
	DisabledAt    *time.Time `gorm:"column:disabled_at" json:"-"`
	LastLogin     *time.Time `gorm:"column:last_login" json:"last_login"`
	Gender        string     `gorm:"column:gender" json:"gender" validate:"oneof=male,female,other" example:"male"`
	Appeal        *string    `gorm:"column:appeal" json:"appeal" validate:"omitempty" example:"I am a Attack Helicopter 🚁"`
}

func (u User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
