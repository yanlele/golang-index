## 结构体与map互转 


### 结构体转为map
```go
package utils

import "reflect"

// 结构体转 map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
```


### map 转为 结构体
```go
package main
 
import (
    "fmt"
    "github.com/goinggo/mapstructure"
)
 
type Person struct {
    Name string
    Age int
}
 
func MapToStruct() {
    mapInstance := make(map[string]interface{})
    mapInstance["Name"] = "liang637210"
    mapInstance["Age"] = 28
 
    var person Person
    //将 map 转换为指定的结构体
    if err := mapstructure.Decode(mapInstance, &person); err != nil {
        fmt.Println(err)
    }
    fmt.Printf("map2struct后得到的 struct 内容为:%v", person)
}
 
func main(){
    MapToStruct()
}
```
