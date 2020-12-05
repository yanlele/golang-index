## 自定义日志输出

### 直接使用内置log包就可以实现日志封装
logging/file.go
```go
package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20200102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filepath string) *os.File {
	/*
		os.Stat：返回文件信息结构描述文件。如果出现错误，会返回*PathError
		type PathError struct {
		    Op   string
		    Path string
		    Err  error
		}
	*/
	_, err := os.Stat(filepath)
	switch {
	/*
		os.IsNotExist：能够接受ErrNotExist、syscall的一些错误，它会返回一个布尔值，能够得知文件不存在或目录不存在
	*/
	case os.IsNotExist(err):
		mkDir()
		break
		/*
			os.IsPermission：能够接受ErrPermission、syscall的一些错误，它会返回一个布尔值，能够得知权限是否满足
		*/
	case os.IsPermission(err):
		log.Fatalf("permission: %v", err)
	}

	/*
		os.OpenFile：调用文件，支持传入文件名称、指定的模式调用文件、文件权限，返回的文件的方法可以用于 I/O。
		如果出现错误，则为*PathError。
		const (
		    // Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
		    O_RDONLY int = syscall.O_RDONLY // 以只读模式打开文件
		    O_WRONLY int = syscall.O_WRONLY // 以只写模式打开文件
		    O_RDWR   int = syscall.O_RDWR   // 以读写模式打开文件
		    // The remaining values may be or'ed in to control behavior.
		    O_APPEND int = syscall.O_APPEND // 在写入时将数据追加到文件中
		    O_CREATE int = syscall.O_CREAT  // 如果不存在，则创建一个新文件
		    O_EXCL   int = syscall.O_EXCL   // 使用O_CREATE时，文件必须不存在
		    O_SYNC   int = syscall.O_SYNC   // 同步IO
		    O_TRUNC  int = syscall.O_TRUNC  // 如果可以，打开时
		)
	*/
	handle, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to openFile : %v", err)
	}
	return handle
}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
```

logging/log.go
```go
package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {


	filePath := getLogFileFullPath()
	// 文件写入器
	F = openLogFile(filePath)

	writers := []io.Writer{
		F,
		os.Stdout, // 控制台标准写入器
	}

	// 多个写入器
	fileAndStdoutWriter := io.MultiWriter(writers...)

	logger = log.New(fileAndStdoutWriter, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
```

### 使用
`logging.Info("info %s", info)`
