package initRouter

import (
	"github.com/gin-gonic/gin"
	"go-index/packages/gin_series/05_db/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("../templates/*")
	} else {
		router.LoadHTMLGlob("packages/gin_series/04_form/templates/*")
	}
	router.StaticFile("/favicon.ico", "packages/gin_series/03_templates/favicon.ico")
	router.Static("/statics", "packages/gin_series/03_templates/statics")
	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}
	// 添加 user
	userRouter := router.Group("/user")
	{
		userRouter.GET("/:name", handler.UserSave)
		userRouter.GET("", handler.UserSaveByQuery)
		userRouter.POST("/register", handler.UserRegister)
	}

	return router
}
