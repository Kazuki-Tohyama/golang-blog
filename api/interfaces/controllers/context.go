package controllers

type Context interface {
	Param(string) string
	PostForm(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
}
