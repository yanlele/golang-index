## go项目配置

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

