package models

import "time"

type User struct {
	ID        int             `json:"id"`
	Name      string          `json:"fullname" gorm:"type: varchar(255)"`
	Email     string          `json:"email" gorm:"type: varchar(255)"`
	Password  string          `json:"password" gorm:"type: varchar(255)"`
	Profile   ProfileResponse `json:"profile"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type UsersProfileResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UsersProductResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UsersCartResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}

func (UsersProductResponse) TableName() string {
	return "users"
}

func (UsersCartResponse) TableName() string {
	return "users"
}
