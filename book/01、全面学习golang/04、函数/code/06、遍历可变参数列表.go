package main

import (
	"bytes"
	"fmt"
)

// 让参数链接为字符串
func joinStrings(slist ...string) string {
	// 定义一个字节缓冲， 用于快速的链接字符串
	var b bytes.Buffer

	for _, s := range slist {
		b.WriteString(s)
	}

	return b.String()
}

func main() {
	fmt.Println(joinStrings("yanle ", "le ", "is ", "coding"))
}
