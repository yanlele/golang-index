package article

import (
	"github.com/gin-gonic/gin"
	"go-index/packages/gin_series/09_restful/model"
	"log"
	"net/http"
	"strconv"
)

func Insert(context *gin.Context) {
	article := model.Article{}
	var id = -1
	if e := context.ShouldBindJSON(&article); e == nil {
		id = article.Insert()
	}
	context.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func GetOne(context *gin.Context) {
	id := context.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		log.Panicln("id 不是 int 类型， id 转换失败", e.Error())
	}
	article := model.Article{
		Id: i,
	}
	art := article.FindById()
	context.JSON(http.StatusOK, gin.H{
		"article": art,
	})
}

func GetAll(context *gin.Context) {
	article := model.Article{}
	articles := article.FindAll()
	context.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}

func DeleteOne(context *gin.Context) {
	id := context.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		log.Panicln("id 不是 int 类型, id 转换失败", e.Error())
	}
	article:=model.Article{Id: i}
	article.DeleteOne()
	context.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

