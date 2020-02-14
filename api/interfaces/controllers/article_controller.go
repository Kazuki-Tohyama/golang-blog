package controllers

import (
	"fmt"
	"strconv"

	"github.com/Kazuki-Tohyama/go-nuxt-blog/api/domain"
	"github.com/Kazuki-Tohyama/go-nuxt-blog/api/interfaces/database"
	"github.com/Kazuki-Tohyama/go-nuxt-blog/api/usecase"
)

type ArticleController struct {
	Interactor usecase.ArticleInteractor
}	

func NewArticleController(sqlHandler database.SqlHandler) *ArticleController {
	return &ArticleController {
		Interactor: usecase.ArticleInteractor {
			ArticleRepository: &database.ArticleRepository {
				SqlHandler: sqlHandler,
			},
		},
	}
}

// POST /article
func (controller *ArticleController) Create(c Context) (err error) {
	a := domain.Article{}
	c.Bind(&a)
	fmt.Println(&a)
	article, err := controller.Interactor.Add(a)
	fmt.Println(article)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, article)
	return
}

// GET /article
func (controller *ArticleController) Index(c Context) (err error) {
	articles, err := controller.Interactor.Articles()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, articles)
	return
}

// GET /article/:id
func (controller *ArticleController) Show(c Context) (err error){
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := controller.Interactor.ArticleById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, article)
	return
}
