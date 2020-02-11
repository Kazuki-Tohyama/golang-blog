package usecase

import (
	"github.com/Kazuki-Tohyama/go-nuxt-blog/api/domain"
)

type ArticleInteractor struct {
	ArticleRepository ArticleRepository
}

func (interactor *ArticleInteractor) Add(a domain.Article) (article domain.Article, err error) {
	article, err = interactor.ArticleRepository.Store(a)
	return
}

func (interactor *ArticleInteractor) Articles() (article domain.Articles, err error) {
	article, err = interactor.ArticleRepository.FindAll()
	return
}

func (interactor *ArticleInteractor) ArticleById(id int) (article domain.Article, err error) {
	article, err = interactor.ArticleRepository.FindById(id)
	return
}
