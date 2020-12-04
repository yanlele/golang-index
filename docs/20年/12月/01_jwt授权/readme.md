## jwt 授权

第一步导入包： `go get -u github.com/dgrijalva/jwt-go`

### 定义创建 token 和 解析 token
```go
package util

import (
	"gin-example/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	
    // 失效时间
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
```

### jwt 中间件
这个中间件用于封装需要jwt验证的接口                 
```go
package jwt

import (
	"gin-example/pkg/e"
	"gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Jwt() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := context.DefaultQuery("token", "")

		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code":    code,
				"message": e.GetMsg(code),
				"data":    data,
			})
			context.Abort()
			return
		}
		context.Next()
	}
}
```

### 如何获取Token
第一步获取用户信息的model
```go
package models

type Auth struct {
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username : username, Password : password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}
```

第二步， 获取auth token 的 api 封装
```go
package api

import (
	"gin-example/models"
	"gin-example/pkg/e"
	"gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(context *gin.Context) {
	username := context.Query("username")
	password := context.Query("password")

	valid := validation.Validation{}
	ok, _ := valid.Valid(&auth{username, password})

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist, _ := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			code = e.ERROR_AUTH_TOKEN

			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    data,
	})
}
```

### 接入router
```go
func InitRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(setting.RunMode)

	router.GET("/auth", api.GetAuth)

	apiv1 := router.Group("/api/v1")
	apiv1.Use(jwt.Jwt())
	{
        ...
	}

	return router
}
```

### 测试
首先通过 get 请求 获取到 auth token, 然后接下来 v1 请求部分， 都带上 token 就可以了。                          
这次实现的是比较简单的部分， 通过 response 和 request query 的方式去鉴权。 正确的方式应该是写入 cookie

### 参考
- [使用 JWT 进行身份校验](https://eddycjy.com/posts/go/gin/2018-02-14-jwt/) 
- [Gin(十二):配合JWT](https://juejin.cn/post/6844903905424310279)
- [Gin(十五):JWT使用(续)](https://juejin.cn/post/6844903982624686088)
