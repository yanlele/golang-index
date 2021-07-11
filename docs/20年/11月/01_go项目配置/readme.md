## go项目配置

<!-- toc -->

- [配置方式一：yaml](#%E9%85%8D%E7%BD%AE%E6%96%B9%E5%BC%8F%E4%B8%80yaml)
- [配置方式二：ini](#%E9%85%8D%E7%BD%AE%E6%96%B9%E5%BC%8F%E4%BA%8Cini)

<!-- tocstop -->

### 配置方式一：yaml

首先装包 `go get -u gopkg.in/yaml.v2`
 
写配置文件： config.yaml
```yaml
addr: :8080

dsn: "root:123456@tcp(127.0.0.1:3306)/gorm_example?charset=utf8&parseTime=True&loc=Local"
max_idle_conn: 100
```

读取配置：config.og
```go
package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Addr 			string		`yaml:"addr"`
	DSN				string		`yaml:"dsn"`
	MaxIdleConn		int			`yaml:"max_idle_conn"`
}

var config *Config

func init() {
	// 加载配置
	err := load("config/config.yaml")
	if err != nil {
		fmt.Println("Failed to load configuration")
		return
	}
}

func load(path string) error {
	result, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(result, &config)
}

func Get() *Config {
	return config
}
```

使用：             
```go
config.Get().DSN
```

优点：简单好用


### 配置方式二：ini
安装依赖包： `go get -u github.com/go-ini/ini`

添加配置文件：app.ini
```ini
# debug or release
RUN_MODE = debug

[app]
PageSize = 10
JwtSecret = 233

RuntimeRootPath = runtime/

ImagePrefixUrl = http://127.0.0.1:8000
ImageSavePath = upload/images/
# MB
ImageMaxSize = 5
ImageAllowExts = .jpg,.jpeg,.png

LogSavePath = logs/
LogSaveName = log
LogFileExt = log
TimeFormat = 20060102

[server]
RunMode = debug
HttpPort = 8080
ReadTimeout = 60
WriteTimeout = 60

[database]
Type = mysql
User = root
Password = 123456
#127.0.0.1:3306
Host = 127.0.0.1:3306
Name = blog
TablePrefix = blog_
```

解析配置文件： setting.go
```go
package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cig          *ini.File
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int
	JwtSecret    string
)

func init() {
	var err error
	Cig, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("加载初始化文件 'conf/app.ini' 文件失败: %v", err)
	}
}

func LoadBase() {
	RunMode = Cig.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cig.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cig.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
```

设置方式2：
把配置直接使用 mapTo 到一个结构体里面
```go
package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

func Setup() {
	appConfig, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("加载初始化文件 'conf/app.ini' 文件失败: %v", err)
	}

	err = appConfig.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("config mapTo AppSetting err : %v", err)
	}
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = appConfig.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("config mapTo ServerSetting err: %v", err)
	}
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = appConfig.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("config mapTo DatabaseSetting err: %v", err)
	}
}
```


使用：                             
方式一
```go
setting.PageSize
```

方式2：                                
```go
sec, err := setting.Cig.GetSection("database")
if err != nil {
    log.Fatalf("Fail to get section 'databse': %v", err)
}
//dbType = sec.Key("TYPE").MustString("mysql")
dbName = sec.Key("NAME").MustString("blog")
user = sec.Key("USER").MustString("root")
password = sec.Key("PASSWORD").String()
host = sec.Key("HOST").String()
tablePrefix = sec.Key("TABLE_PREFIX").String()
```

方式3：
```go
dbName = setting.DatabaseSetting.Name
```


### 配置方式三：viper
首先需要本地装包： `go get github.com/spf13/viper`

#### 最基础的使用
```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
)

func main() {
	// 获取项目目录
	workDir, _ := os.Getwd()
	
	v := viper.New()
	v.SetConfigFile(path.Join(workDir, "src/config.yaml"))
	if err := v.ReadInConfig(); err != nil {
		fmt.Println("配置文件读取失败: ", err)
		return
	}
	fmt.Println(v.Get("name"))
}
```

其中 `src/config.yaml` 文件配置如下：
```yaml
name: "test"
```

#### 使用结构体
```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
)

type ServerConfig struct {
	Name string `json:"name"`
}

func main() {
	// 获取项目目录
	workDir, _ := os.Getwd()

	v := viper.New()
	v.SetConfigFile(path.Join(workDir, "src/config.yaml"))
	if err := v.ReadInConfig(); err != nil {
		fmt.Println("配置文件读取失败: ", err)
		return
	}

	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		fmt.Println("解析结构体失败", err)
		return
	}

	fmt.Println("结构体： ", serverConfig.Name)
	fmt.Println(v.Get("name"))
}
```

#### 使用环境变量来读取不同的配置
获取环境变量的方式：
```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// 这两种方式都可以获取环境变量
func getEnvInfo(env string) interface{} {
	viper.AutomaticEnv()
	return viper.Get(env)
}

// 这两种方式都可以获取环境变量
func getEnv(env string) string  {
	return os.Getenv(env)
}

func main() {
	fmt.Println("getEnvInfo: ", getEnvInfo("IS_DEV"))
	fmt.Println("getEnv: ", getEnv("IS_DEV"))
}
```
运行： 
```
$ IS_DEV=123 go run main.go
getEnvInfo:  123
getEnv:  123
```


### 参考文档
- [go项目中环境变量的配置](https://juejin.cn/post/6983290445577060382)

