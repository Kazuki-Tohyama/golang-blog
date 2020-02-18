package usecase

import (
	"github.com/Kazuki-Tohyama/go-nuxt-blog/api/domain"
)

type ArticleRepository interface {
	StoreArticle(domain.Article) (domain.Article, error)
	FindArticleByUserId(int) (domain.Article, error)
	FindAllArticles() (domain.Articles, error)
}
