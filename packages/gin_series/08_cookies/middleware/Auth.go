package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-index/packages/gin_series/12_jwt/config"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		println("已经授权")
		cookie, e := context.Request.Cookie("user_cookie")
		if e == nil {
			context.SetCookie(cookie.Name, cookie.Value, 1000, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
			context.Next()
		} else {
			context.Abort()
			context.HTML(http.StatusUnauthorized, "401.tmpl", nil)
		}
	}
}

// 验证token
func parseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(
		token,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (i interface{}, e error) {
			return []byte(config.Secret), nil
		})

	if err == nil && jwtToken != nil {
		claim, ok := jwtToken.Claims.(*jwt.StandardClaims)
		if ok && jwtToken.Valid {
			return claim, err
		}
	}
	return nil, err
}
