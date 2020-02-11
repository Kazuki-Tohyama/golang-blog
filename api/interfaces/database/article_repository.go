package database

import (
	"github.com/Kazuki-Tohyama/go-nuxt-blog/api/domain"
)

type ArticleRepository struct {
	SqlHandler
}

func (repo *ArticleRepository) Store(a domain.Article) (article domain.Article, err error) {
	if err = repo.Create(&a).Error; err != nil {
		return
	}
	article = a
	return
}

func (repo *ArticleRepository) FindById(id int) (article domain.Article, err error) {
	if err = repo.Find(&article, id).Error; err != nil {
		return
	}
	return
}

func (repo *ArticleRepository) FindAll() (articles domain.Articles, err error) {
	if err = repo.Find(&articles).Error; err != nil {
		return
	}
	return
}

