package _0_测试demo

import "testing"

func BenchmarkGetArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetArea(40, 50)
	}
}
