package handler

import (
	"github.com/gin-gonic/gin"
	"go-index/packages/gin_series/05_db/model"
	"log"
	"net/http"
)

func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户"+username+"已经保存")
}

// 通过 query 方法进行获取参数
func UserSaveByQuery(context *gin.Context) {
	username := context.Query("name")
	//age := context.Query("age")
	age := context.DefaultQuery("age", "20")
	context.String(http.StatusOK, "用户:"+username+",年龄:"+age+"已经保存")
}

func UserRegister(context *gin.Context) {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		log.Println("err ->", err.Error())
		context.String(http.StatusBadRequest, "输入的数据不合法")
	}
	log.Println("email", user.Email, "password", user.Password, "password again", user.PasswordAgain)
	context.Redirect(http.StatusMovedPermanently, "/")
}
