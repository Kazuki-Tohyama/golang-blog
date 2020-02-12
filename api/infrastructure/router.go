package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/Kazuki-Tohyama/go-nuxt-blog/api/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	articleController := controllers.NewArticleController(NewSqlHandler())

	router.GET("/articles", func(c *gin.Context) { articleController.Index(c) })
	router.POST("/articles", func(c *gin.Context) { articleController.Create(c) })
	router.GET("/articles/:id", func(c *gin.Context) { articleController.Show(c) })

	Router = router

	Router.Run(":8080")
}
