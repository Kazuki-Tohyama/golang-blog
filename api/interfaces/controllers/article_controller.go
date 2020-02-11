package controllers

import (
	"github.com/Kazuki-Tohyama/go-nuxt-blog/api/domain"
	"github.com/Kazuki-Tohyama/go-nuxt-blog/api/interfaces/database"
	"github.com/Kazuki-Tohyama/go-nuxt-blog/api/usecase"
	"strconv"
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

// POSt /article
func (controller *ArticleController) Create(c Context) {
	a := domain.Article{}
	c.Bind(&a)
	article, err := controller.Interactor.Add(a)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, article)
}

// GET /article
func (controller *ArticleController) Index(c Context) {
	articles, err := controller.Interactor.Articles()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, articles)
}

// GET /article/:id
func (controller *ArticleController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := controller.Interactor.ArticleById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, article)
}
