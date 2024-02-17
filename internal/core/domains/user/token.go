package user_domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Token struct {
	UserUUID  uuid.UUID
	NewEmail  string
	OldMail   string
	FirstName string
	LastName  string
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
}

func (u Token) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *Token) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
