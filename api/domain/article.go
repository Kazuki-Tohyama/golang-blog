package domain

type Article struct {
	ID int `gorm:"primary_key"`
	Title string
	Body string	
}

type Articles []Article 

