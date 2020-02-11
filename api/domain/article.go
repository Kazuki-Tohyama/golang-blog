package domain

type Article struct {
	ID int
	Title string
	Body string	
	tag []string
}

type Articles []Article 

