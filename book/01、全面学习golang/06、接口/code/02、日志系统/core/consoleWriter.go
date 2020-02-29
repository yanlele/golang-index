package core
import (
	"fmt"
	"os"
)

type consoleWriter struct {
}

// 实现LogWriter的Write()方法
func (f *consoleWriter) Write(data interface{}) error {
	str := fmt.Sprintf("%v\n", data)
	_, err := os.Stdout.Write([]byte(str))
	return err
}

// 创建民领航写入器实例
func NewConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}
