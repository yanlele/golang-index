package main

import "fmt"

var address = "重庆"

// 值引用
// 这样的引用是不会改变原始值的数据
func setAddress(address string) string {
	address = "成都"
	return address
}


// 地址引用
// 这样的引用是会更改原始值的数据的
func setAddressPoint(address *string) string {
	*address = "成都"
	return *address
}

func main() {
	//setAddressPoint(&address)
	fmt.Println(setAddress(address))
	fmt.Println(address)
}
