package initRouter

import (
	"github.com/gin-gonic/gin"
	"go-index/packages/gin_series/08_cookies/middleware"
	"go-index/packages/gin_series/12_jwt/handler/article"
	"net/http"
	"time"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", middleware.Auth(), func(context *gin.Context) {
		context.JSON(http.StatusOK, time.Now().Unix())
	})
	articleRouter := router.Group("")
	{
		articleRouter.GET("/article/:id", article.GetOne)
		articleRouter.GET("/articles", article.GetAll)
		articleRouter.POST("/article", article.Insert)
		articleRouter.DELETE("/article/:id", article.DeleteOne)
	}
	return router
}
