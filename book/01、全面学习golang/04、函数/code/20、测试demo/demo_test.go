package _0_测试demo

import "testing"

func TestGetArea(t *testing.T) {
	area := GetArea(40, 50)

	if area != 2000 {
		t.Error("测试失败")
	}
}
