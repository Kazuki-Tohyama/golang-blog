package infrastructure

import (
	"github.com/Kazuki-Tohyama/go-nuxt-blog/api/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	articleController := controllers.NewArticleController(NewSqlHandler())

	router.GET("/articles", func(c *gin.Context) { articleController.Index(c) })
	router.POST("/articles", func(c *gin.Context) { articleController.Create(c) })
	router.GET("/articles/:id", func(c *gin.Context) { articleController.Show(c) })

	router.Run(":8080")
}
