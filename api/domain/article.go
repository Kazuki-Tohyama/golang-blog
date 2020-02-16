package domain

import (
	"time"
)

type Article struct {
	ID int `gorm:"column:id;primary_key"`
	Title string `gorm:"column:title"`
	Body string	`gorm:"column:body;type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Articles []Article 

