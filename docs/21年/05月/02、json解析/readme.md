## json 解析


### 一个对象编码为 json 数据对象
`func Marshal(v interface{}) ([]byte, error)`
```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func main() {
	group := ColorGroup{
		ID:     1,
		Name:   "reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	// func Marshal(v interface{}) ([]byte, error)
	b, err := json.Marshal(group)

	if err != nil {
		fmt.Println("error: ", err)
	}

	// 输出方式1
	fmt.Println(string(b))

	// 实处方式2
	os.Stdout.Write(b)
}
```

### json数据解码
`func Unmarshal(data []byte, v interface{}) error`                          
```go
package main

import (
	"encoding/json"
	"fmt"
)

var jsonBlob = []byte(`[
{"Name": "Platypus", "Order": "Monotremata"},
{"Name": "Quoll", "Order": "Dasyuromorphia"}
]`)

type Animal struct {
	Name  string
	Order string
}

func main() {
	var animals []Animal

	// func Unmarshal(data []byte, v interface{}) error
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Printf("%+v", animals)
}
```


### 手动配置结构体的成员和JSON字段的对应关系
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string `json:"msg_name"`       // 对应JSON的msg_name
	Body string `json:"body,omitempty"` // 如果为空置则忽略字段
	Time int64  `json:"-"`              // 直接忽略字段
}

func main() {
	var m = Message{
		Name: "Alice",
		Body: "",
		Time: 1294706395881547000,
	}

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(string(data))

}
```

### 参考文章
[https://www.jianshu.com/p/d4a66eaa46d2](https://www.jianshu.com/p/d4a66eaa46d2)
