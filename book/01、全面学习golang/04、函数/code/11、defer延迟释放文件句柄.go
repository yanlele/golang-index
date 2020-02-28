package main

import "os"

/*
正常的代码是这样子的
func fileSize(filename string) int64 {
	f, err := os.Open(filename)

	if err != nil {
		return 0
	}

	info, err := f.Stat()

	if err != nil {
		_ = f.Close()
		return 0
	}

	// 获取文件大小
	size := info.Size()

	_ = f.Close()

	return size
}
*/

/*
用deffer 修改之后的代码是这样子的
*/

func fileSize(filename string) int64 {
	f, err := os.Open(filename)

	if err != nil {
		return 0
	}

	defer f.Close()

	info, err := f.Stat()

	if err != nil {
		return 0
	}

	size := info.Size()

	return size
}
