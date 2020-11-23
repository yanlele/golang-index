package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-index/packages/gin_series/12_jwt/config"
	"go-index/packages/gin_series/12_jwt/model"
	"net/http"
	"time"
)

func CreateJwt(context *gin.Context) {
	// 获取用户信息
	user := &model.User{}
	result := &model.Result{
		Code:    http.StatusOK,
		Message: "登录成功",
		Data:    nil,
	}
	e := context.BindJSON(&user)
	if e != nil {
		result.Message = "数据绑定失败"
		result.Code = http.StatusUnauthorized
		context.JSON(http.StatusUnauthorized, gin.H{
			"result": result,
		})
	}

	u := user.QueryByUsername()
	if u.Password == user.Password {
		expiresTime := time.Now().Unix() + int64(config.OndDayOfHours)
		claims := jwt.StandardClaims{
			Audience:  user.Username,     // 受众
			ExpiresAt: expiresTime,       // 失效时间
			Id:        string(user.ID),   // 编号
			IssuedAt:  time.Now().Unix(), // 签发时间
			Issuer:    "gin hello",       // 签发人
			NotBefore: time.Now().Unix(), // 生效时间
			Subject:   "login",           // 主题
		}
		var jwtSecret = []byte(config.Secret)
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, err := tokenClaims.SignedString(jwtSecret)
		if err == nil {
			result.Message = "登录成功"
			result.Data = "bearer " + token
			result.Code = http.StatusOK
			context.JSON(http.StatusOK, gin.H{
				"result": result,
			})
		} else {
			result.Message = "登录失败"
			result.Code = http.StatusOK
			context.JSON(result.Code, gin.H{
				"result": result,
			})
		}
	} else {
		result.Message = "登录失败"
		result.Code = http.StatusOK
		context.JSON(result.Code, gin.H{
			"result": result,
		})
	}
}

func Register(context *gin.Context) {
	user := model.User{}
	result := &model.Result{
		Code:    http.StatusOK,
		Message: "登录成功",
		Data:    nil,
	}
	e := context.BindJSON(&user)
	if e != nil {
		result.Message = "绑定数据失败"
		result.Code = http.StatusUnauthorized
		context.JSON(http.StatusUnauthorized, gin.H{
			"result": result,
		})
	}
	if user.Insert() {
		result.Message = "注册成功"
		result.Code = http.StatusOK
		context.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	} else {
		result.Message = "注册失败"
		result.Code = http.StatusOK
		context.JSON(http.StatusOK, gin.H{
			"result": result
		})
	}
}
