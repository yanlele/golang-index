package initRouter

import (
	"github.com/gin-gonic/gin"
	"go-index/packages/gin_series/12_jwt/handler/article"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	articleRouter := router.Group("")
	{
		articleRouter.GET("/article/:id", article.GetOne)
		articleRouter.GET("/articles", article.GetAll)
		articleRouter.POST("/article", article.Insert)
		articleRouter.DELETE("/article/:id", article.DeleteOne)
	}
	return router
}
