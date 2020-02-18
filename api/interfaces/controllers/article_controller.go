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
func (controller *ArticleController) CreateArticle(c Context) {
	uid := c.PostForm("uid")
	userId, uidErr := strconv.Atoi(uid)
	a := domain.Article{}
	a.UserID = userId
	c.Bind(&a)
	fmt.Println(&a)
	article, err := controller.Interactor.AddArticle(a)
	fmt.Println(article)
	if uidErr != nil || err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, article)
	return
}

// GET /article
func (controller *ArticleController) IndexArticles(c Context) {
	articles, err := controller.Interactor.Articles()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, articles)
	return
}

// GET /article/:id
func (controller *ArticleController) ShowArticles(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := controller.Interactor.ArticleByUserId(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, article)
	return
}
