package main

import (
	"fmt"
	"strings"
)

func StringProcess(list []string, chain []func(string) string) {
	// 遍历每个字符串
	for index, str := range list {
		// 需要处理的字符串
		result := str

		// 遍历处理链
		for _, proc := range chain {

			// 输入一个字符串进行处理， 返回数据作为下一个处理链的输入
			result = proc(result)
		}

		// 结果返回给切片
		list[index] = result
	}
}

// 移除前缀
func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func main() {
	// 待处理的字符串列表
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	// 处理函数作用连
	chain := []func(string) string{
		removePrefix,      // 移除前缀
		strings.TrimSpace, // 移除空格
		strings.ToUpper,   // 转大写
	}

	// 处理字符串
	StringProcess(list, chain)

	for _, str := range list {
		fmt.Println(str)
	}
}
