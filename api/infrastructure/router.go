package infrastructure

import (
	"github.com/Kazuki-Tohyama/go-nuxt-blog/api/interfaces/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	controller := controllers.NewArticleController(NewSqlHandler())

	router.GET("/articles", func(c *gin.Context) { controller.IndexArticles(c) })
	router.POST("/articles", func(c *gin.Context) { controller.CreateArticle(c) })
	router.GET("/articles/:id", func(c *gin.Context) { controller.ShowArticles(c) })

	router.Run(":8080")
}
