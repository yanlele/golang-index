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
