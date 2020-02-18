package domain

import (
	"time"
)

type User struct {
	ID int `gorm:"column:id;primary_key"`
	UserName string `gorm:"column:user_name;size:255"`
	Password string `gorm:"column:password"`
	Articles []Article
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Users []User

