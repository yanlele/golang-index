## 文件上传

### 先定义文件处理
```go
package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

/* 获取文件大小 */
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)
	return len(content), err
}

func GetExt(filename string) string {
	return path.Ext(filename)
}

func CheckNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}
	return nil
}

func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
```

### 文件名处理
一般不会直接将上传的图片名暴露出来，因此我们对图片名进行 MD5 来达到这个效果
```go
package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
```

### 封装 image 的处理逻辑
```go
package upload

import (
	"fmt"
	"gin-example/pkg/file"
	"gin-example/pkg/logging"
	"gin-example/pkg/setting"
	"gin-example/pkg/util"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

func GetImageFullUrl(name string) string {
	return setting.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}

func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		logging.Warn(err)
	}
	return size <= setting.AppSetting.ImageMaxSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err : %v", err)
	}
	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.isNotExistMkDir err :%v", err)
	}
	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src : %v", src)
	}
	return nil
}
```

### 编写上传图片的业务逻辑
在 routers/api 目录下 新建 upload.go 文件，写入文件内容:
```go
package api

import (
	"gin-example/pkg/e"
	"gin-example/pkg/logging"
	"gin-example/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadImage(context *gin.Context) {
	code := e.SUCCESS
	resData := make(map[string]interface{})
	file, image, err := context.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		code = e.ERROR
		context.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": e.GetMsg(code),
			"data":    resData,
		})
		return
	}
	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()

		src := fullPath + imageName
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if err := context.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				resData["image_url"] = upload.GetImageFullUrl(imageName)
				resData["image_save_url"] = savePath + imageName
			}
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    resData,
	})
}
```


### 封装api
打开 routers/router.go 文件，增加路由 r.POST("/upload", api.UploadImage)
```go
func InitRouter() *gin.Engine {
	r := gin.New()
    ...
	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		...
	}

	return r
}
```

### 访问上传的文件
这个时候其实已经能够上传文件了， 但是还是无法访问到文件。
打开 routers/router.go 文件，
增加路由 r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))，如：
```go
func InitRouter() *gin.Engine {
    ...
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)
    ...
}
```

这个时候访问文件， 就可以访问到了。

