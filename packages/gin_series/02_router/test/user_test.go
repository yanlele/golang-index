package test

import (
	"github.com/stretchr/testify/assert"
	"go-index/packages/gin_series/02_router/initRouter"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestUserSave(t *testing.T) {
	username := "yanle"
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户："+username+"已经保存", w.Body.String())
}

func TestUserSaveQuery(t *testing.T) {
	username := "yanle"
	age := 18
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user?name="+username+"&age="+strconv.Itoa(age), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户:"+username+",年龄:"+strconv.Itoa(age)+"已经保存", w.Body.String())
}

func TestUserSaveWithNotAge(t *testing.T) {
	username := "yanle"
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user?name="+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户:"+username+",年龄:20已经保存", w.Body.String())
}
