## 优雅的重启服务

### 怎样算优雅

#### 目的

- 不关闭现有连接（正在运行中的程序）
- 新的进程启动并替代旧进程
- 新的进程接管新的连接
- 连接要随时响应用户的请求，当用户仍在请求旧进程时要保持连接，新用户应请求新进程，不可以出现拒绝请求的情况


#### 流程

1、替换可执行文件或修改配置文件

2、发送信号量 SIGHUP

3、拒绝新连接请求旧进程，但要保证已有连接正常

4、启动新的子进程

5、新的子进程开始 Accet

6、系统将新的请求转交新的子进程

7、旧进程处理完所有旧连接后正常结束


### 实现
我们借助 `fvbock/endless` 来实现 Golang HTTP/HTTPS 服务重新启动的零停机

安装： `go get -u github.com/fvbock/endless`

修改main.go
```go
package main

import (
    "fmt"
    "log"
    "syscall"

    "github.com/fvbock/endless"

    "gin-blog/routers"
    "gin-blog/pkg/setting"
)

func main() {
    endless.DefaultReadTimeOut = setting.ReadTimeout
    endless.DefaultWriteTimeOut = setting.WriteTimeout
    endless.DefaultMaxHeaderBytes = 1 << 20
    endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

    server := endless.NewServer(endPoint, routers.InitRouter())
    server.BeforeBegin = func(add string) {
        log.Printf("Actual pid is %d", syscall.Getpid())
    }

    err := server.ListenAndServe()
    if err != nil {
        log.Printf("Server err: %v", err)
    }
}
```

启动时候有一个 pid, 之后需要暂停服务做修改的时候， 直接 在另外一个终端执行 `kill -1 pid` ，检验先前服务的终端效果

可以看到该命令已经挂起，并且 fork 了新的子进程 pid 

这时候在 postman 上再次访问我们的接口，你可以惊喜的发现，他“复活”了！

这就完成了一次正向的流转了

你想想，每次更新发布、或者修改配置文件等，只需要给该进程发送SIGTERM 信号，而不需要强制结束应用，是多么便捷又安全的事！


具体使用请看下面的文章就行了
[优雅的重启服务](https://eddycjy.com/posts/go/gin/2018-03-15-reload-http/)


