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
