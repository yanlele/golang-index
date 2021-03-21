# 参数验证

## github.com/astaxie/beego/validation


### 装包
` go get -u github.com/astaxie/beego/validation`

### 使用
```go
// 验证
valid := validation.Validation{}
valid.Required(name, "name").Message("名称不能为空")
valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
valid.Required(createdBy, "created_by").Message("创建人不能为空")
valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
valid.Range(state, 0, 1, "state").Message("状态只允许0或1")


code := e.INVALID_PARAMS
if !valid.HasErrors() {
    if exist, _ := models.ExistTagByName(name); !exist {
        code = e.SUCCESS
        _ = models.AddTag(name, state, createdBy)
    } else {
        code = e.ERROR_EXIST_TAG
    }
}
```


## github.com/asaskevich/govalidator
### 装包
`go get github.com/asaskevich/govalidator`

### 使用1
```go
package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"strings"
)

type addressStruct struct {
	AddressName string
	AddressCode string
}

type userStruct struct {
	Name    string        `valid:"required~缺少姓名"`
	Age     int           `valid:"required~缺少年龄"`
	Address addressStruct `valid:"required~缺少地址信息"`
}

func main() {
	user := userStruct{
		Name: "yanle",
		Age: 30,
		Address: addressStruct{"12", ""},
	}

	result, err := govalidator.ValidateStruct(user)
	fmt.Println("result: ", result)
	if err != nil {
		fmt.Println(strings.Split(err.Error(), ";")[0])
	}
}
```

### 使用2
[api 列表](https://github.com/asaskevich/govalidator#list-of-functions)



