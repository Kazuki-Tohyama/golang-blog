package usecase

import (
	"github.com/Kazuki-Tohyama/go-nuxt-blog/api/domain"
)

type ArticleRepository interface {
	Store(domain.Article) (domain.Article, error)
	FindById(int) (domain.Article, error)
	FindAll() (domain.Articles, error)
}
