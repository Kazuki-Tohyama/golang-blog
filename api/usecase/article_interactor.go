package usecase

import (
	"github.com/Kazuki-Tohyama/go-nuxt-blog/api/domain"
)

type ArticleInteractor struct {
	ArticleRepository ArticleRepository
}

func (interactor *ArticleInteractor) AddArticle(a domain.Article) (article domain.Article, err error) {
	article, err = interactor.ArticleRepository.StoreArticle(a)
	return
}

func (interactor *ArticleInteractor) Articles() (article domain.Articles, err error) {
	article, err = interactor.ArticleRepository.FindAllArticles()
	return
}

func (interactor *ArticleInteractor) ArticleByUserId(id int) (article domain.Article, err error) {
	article, err = interactor.ArticleRepository.FindArticleByUserId(id)
	return
}
