package models

import "time"

type Transaction struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Phone     string    `json:"phone" gorm:"type: varchar(255)"`
	Address   string    `json:"address" gorm:"type: text"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
