package main

import (
	"encoding/json"
	"fmt"
)

type Screen struct {
	Size       float64
	ResX, ResY int
}

type Battery struct {
	Capacity int
}

func genJsonData() []byte {
	// 匿名结构体
	raw := &struct {
		Screen
		Battery
		HasTouchId bool
	}{
		Screen: Screen{
			Size: 5.5,
			ResX: 1902,
			ResY: 1080,
		},
		Battery: Battery{
			Capacity: 2910,
		},
		HasTouchId: true,
	}

	jsonData, _ := json.Marshal(raw)
	return jsonData
}

func main() {
	jsonData := genJsonData()
	fmt.Println(string(jsonData))

	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}

	// 反序列化到 screenAndTouch
	_ = json.Unmarshal(jsonData, &screenAndTouch)

	// 输出 screenAndTouch 的详细结构
	fmt.Printf("%+v\n", screenAndTouch)

	batteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}

	// 反序列化到batteryAndTouch
	_ = json.Unmarshal(jsonData, &batteryAndTouch)

	// 输出screenAndTouch的详细结构
	fmt.Printf("%+v\n", batteryAndTouch)

}
