package user_domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID          *uuid.UUID `json:"uuid,omitempty"`
	Password      string     `json:"password,omitempty"`
	FirstName     string     `json:"firstname" validate:"required,alphaunicode" example:"Amine Beren"`
	LastName      string     `json:"lastname" validate:"required,alphaunicode" example:"Çelik"`
	Lang          string     `json:"lang" validate:"required,oneof=en tr" example:"tr"` //🤘
	Email         string     `json:"email" validate:"omitempty,email" example:"amine@un-defined.com"`
	EmailVerified bool       `json:"email_verified"`
	MasterAdmin   bool       `json:"-"`
	Banned        bool       `json:"-"`
	MFAEnabled    bool       `json:"mfa_enabled"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"-"`
	Disabled      bool       `json:"-"`
	DisabledAt    *time.Time `json:"-"`
	LastLogin     *time.Time `json:"last_login"`
	Gender        string     `json:"gender"`
	Appeal        string     `json:"appeal"`
}

func (u User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
