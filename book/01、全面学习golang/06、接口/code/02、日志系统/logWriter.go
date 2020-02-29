package main

type LogWriter interface {
	Write(data interface{}) error
}

type Logger struct {
	writerList []LogWriter
}

// 注册一个日志写入器
func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writerList = append(l.writerList, writer)
}

// 将一个data类型的输入写入日志
func (l *Logger) Log(data interface{}) {
	for _, writer := range l.writerList {
		_ = writer.Write(data)
	}
}

func newLogger() *Logger {
	return &Logger{}
}
