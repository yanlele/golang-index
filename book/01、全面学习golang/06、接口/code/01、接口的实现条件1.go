package main

import "fmt"

type DataWriter interface {
	WriteData(data interface{}) error
}

type file struct {
}

// 实现 DataWriter 接口的 WriteData 方法
func (d *file) WriteData(data interface{}) error {
	fmt.Println("WriteData: ", data)
	return nil
}

func main() {
	// 实例化file
	f := new(file)

	// 申明一个DataWriter接口
	var writer DataWriter
	writer = f
	_ = writer.WriteData("data")
}
