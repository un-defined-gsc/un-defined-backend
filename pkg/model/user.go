package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Gender    string    `json:"gender"`
	Appeal    string    `json:"appeal"`
	Comment   []Comment `gorm:"foreignKey:UserID"`
	Like      []Like    `gorm:"foreignKey:UserID"`
	Image     []Image   `gorm:"foreignKey:UserID"`
	Post      []Post    `gorm:"foreignKey:UserID"`
	Tag       []Tag     `gorm:"foreignKey:UserID"`
}

type UserDTO struct {
	ID       uuid.UUID `gorm:"primary_key" json:"id"`
	Name     string    `json:"name"`
	Surname  string    `json:"surname"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Gender   string    `json:"gender"`
	Appeal   string    `json:"appeal"`
}

func ToUser(usersDTO *UserDTO) *User {
	return &User{
		Name:     usersDTO.Name,
		Surname:  usersDTO.Surname,
		Email:    usersDTO.Email,
		Password: usersDTO.Password,
		Gender:   usersDTO.Gender,
		Appeal:   usersDTO.Appeal,
	}
}

func ToUserDTO(users *User) *UserDTO {
	return &UserDTO{
		ID:       users.ID,
		Name:     users.Name,
		Surname:  users.Surname,
		Email:    users.Email,
		Password: users.Password,
		Gender:   users.Gender,
		Appeal:   users.Appeal,
	}
}
