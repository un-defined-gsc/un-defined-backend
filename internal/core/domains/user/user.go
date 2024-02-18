package user_domain

import (
	"encoding/json"
	"time"

	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
)

type User struct {
	base_domain.Base
	Password      string     `gorm:"column:password;type:TEXT;NOT NULL" json:"password,omitempty"`
	FirstName     string     `gorm:"column:first_name;type:TEXT;NOT NULL" json:"firstname" validate:"required,alphaunicode" example:"Resul"`
	LastName      string     `gorm:"column:last_name;type:TEXT;NOT NULL" json:"lastname" validate:"required,alphaunicode" example:"Çelik"`
	Lang          string     `gorm:"column:lang;type:TEXT;default:en;NOT NULL" json:"lang" validate:"required,oneof=en tr" example:"tr"` //🤘
	Email         string     `gorm:"column:email;unique;type:TEXT;NOT NULL" json:"email" validate:"omitempty,email" example:"resul@mon.time"`
	EmailVerified bool       `gorm:"column:email_verified;type:BOOLEAN;NOT NULL;default:FALSE" json:"email_verified"`
	MasterAdmin   bool       `gorm:"column:master_admin;type:BOOLEAN;NOT NULL;default:FALSE" json:"-"`
	Banned        bool       `gorm:"column:banned;type:BOOLEAN;NOT NULL;default:FALSE" json:"-"`
	MFAEnabled    bool       `gorm:"column:mfa_enabled;type:BOOLEAN;NOT NULL;default:FALSE" json:"mfa_enabled"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"-"`
	Disabled      bool       `gorm:"column:disabled;type:BOOLEAN;NOT NULL;default:FALSE" json:"-"`
	DisabledAt    *time.Time `gorm:"column:disabled_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"-"`
	LastLogin     *time.Time `gorm:"column:last_login;type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"last_login"`
	Gender        string     `gorm:"column:gender;type:TEXT;NOT NULL" json:"gender" validate:"required,oneof=male,female,other" example:"male"`
	Appeal        *string    `gorm:"column:appeal;type:TEXT" json:"appeal" validate:"omitempty" example:"I am a Attack Helicopter 🚁"`
}

func (u User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

func (User) TableName() string {
	return "t_users"
}
