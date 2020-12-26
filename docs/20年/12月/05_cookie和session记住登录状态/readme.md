## cookie和session记住登录状态

其实这个用 jwt 就可以实现了 emmm

如果不想用 jwt , 还可以这样

### 安装必要的安装包
`go get -u github.com/gin-contrib/sessions`

### 具体实现如下

定义一个 auth.go 文件， 做 auth 相关信息的暂存
```go
package controllers

import "github.com/gin-gonic/gin"

type Auth struct {
	Id         int
	Name       string
	Avatar     string
	Profession string
}

func (a Auth) GetAuth(c *gin.Context) Auth {
	auth, exists := c.Get("auth")
	if !exists {
		auth = Auth{
			Id:         0,
			Name:       "",
			Avatar:     "/static/logoh.png",
			Profession: "",
		}
	}
	return auth.(Auth)
}
```

中间件：
```go
package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"le-blog/controllers"
	"le-blog/utils"
)

// 注入 auth
func SetAuth(c *gin.Context) {
	sess := sessions.Default(c)
	auth := sess.Get("auth")

	if auth != nil {
		c.Set("auth", auth)
	}
	c.Next()
}

func Authorization(c *gin.Context) {
	auth := controllers.Auth{}.GetAuth(c)
	if auth.Id == 0 {
		// 用户没有登录的情况
		utils.Redirect(c, "/login")
		return
	}
	c.Next()
}
```


在 router 初始化的时候 routerInit.go
```go
func Init() *gin.Engin {
    app := gin.Default()
    
    // 完全不知道这个有啥用？
    gob.Register(controllers.Auth{})

    // 添加cookie 和 session
    store := cookie.NewStore([]byte("secret"))
    app.Use(sessions.Sessions(CookieSessionKey, store))

	// 设置用户信息
	app.Use(middleware.SetAuth)
    // ......

    routers.Api(app)
    routers.Home(app)
    return app
}
```

在路由注册的时候， 例如 home.go                        
需要使用权限 auth 的地方直接添加 middleware 就可以
```go
package routers

import (
	"github.com/gin-gonic/gin"
	"le-blog/controllers"
	"le-blog/middleware"
)

func Home(r *gin.Engine) {
	home := r.Group("/")
	{
		// 首页
		home.GET("/", controllers.Index)

		article := home.Group("/article", middleware.Authorization)
		{
			article.GET("/user", controllers.UserArticleList)
			article.GET("/create", controllers.CreateArticle)
			article.POST("/create", controllers.SaveArticle)
			article.GET("/edit/:id", controllers.EditArticle)
			article.GET("/delete/:id", controllers.DelArticle)
		}

		// 个人中心
		user := home.Group("/user", middleware.Authorization)
		{
			user.GET("/update_pwd", controllers.UpdatePwd)
			user.POST("/update_pwd", controllers.DoUpdatePwd)
		}

		// 文章详情
		home.GET("/detail/:id", controllers.Detail)

		// 标签页面
		tag := home.Group("/tags")
		{
			tag.GET("/", controllers.TagIndex)
			tag.GET("/title/:name", controllers.GetArticleByTagName)
			tag.GET("/ajax/list", controllers.AjaxTags)
			tag.POST("/add", controllers.AddTags)
		}

		home.GET("/archives", controllers.Archives)

		// 注册
		home.GET("/join", controllers.Register)
		home.POST("/join", controllers.DoRegister)

		// sign in
		home.GET("/login", controllers.Login)
		home.POST("/login", controllers.DoLogin)
		home.GET("/logout", controllers.Logout)
	}
}
```

写入 auth， 例如在登录的时候
```go
func DoLogin(c *gin.Context) {
    // ......
    auth := &Auth{
		Id:         int(user.ID),
		Name:       user.Name,
		Avatar:     user.Avatar,
		Profession: user.Profession,
	}

	session := sessions.Default(c)
	session.Set("auth", auth)
	err = session.Save()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, (&utils.Response{
		Status: 0,
		Data:   nil,
		Msg:    "",
	}).SuccessResponse())
	return
}
```

清除 auth, 例如在登出的时候
```go
// method : GET
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		panic(err)
	}
	c.Redirect(http.StatusFound, "/")
	return
}
```

需要获取到 auth 的时候
```go
auth := Auth{}.GetAuth(c)
```
