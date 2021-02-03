# 参数验证

## 方式1：


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
