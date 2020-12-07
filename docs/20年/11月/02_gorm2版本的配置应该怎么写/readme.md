## gorm2配置应该怎么写


### 链接的配置
```go
package models

import (
	"database/sql"
	"fmt"
	"gin-example/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

/* 初始化数据库链接 */
var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err error
		//dbType,
		dbName, user, password, host, tablePrefix string
	)
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

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)

	// 日志相关
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // 禁用彩色打印
		},
	)

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,  // string 类型字段的默认长度
		DisableDatetimePrecision:  true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix, // 所有 table 前缀
			SingularTable: true,        // 最后 table 不加s
		},
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	//gorm.DefaultTableNameHandler = func() {}
	var sqlDB *sql.DB
	sqlDB, err = db.DB()
	if err != nil {
		log.Fatalf("get db.BD() error: %v", err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	defer sqlDB.Close()
}
```

### 初始化数据库的时候创建表
使用方式： `DB.AutoMigrate(&models.AdminUser{})`

demo:               
```go
package database

import (
	"go-gorm-example/config"
	"go-gorm-example/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.Get().DSN,
		DefaultStringSize:         256,  // string 类型字段的默认长度
		DisableDatetimePrecision:  true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		log.Panicln("链接错误")
	} else {
		err := DB.AutoMigrate(&models.AdminUser{})
		if err != nil {
			log.Println("创建表失败")
		}
		log.Println("链接成功")
	}
}
```

### model 更改链接table名                                
`TableName() string`                                    
```go
package model

type Student struct {
	Id uint `json:"id"`
	Number uint `json:"number"`
	Name string `json:"name"`
	Gender string `gorm:"type:enum('1', '2');default:'1'"`
	Phone string `json:"phone"`
	Age int `json:"age"`
	ClassNumber string `gorm:"class_number;default:null"`
	Email string `json:"email"`
	Address string `json:"address"`
	Birthday int64 `gorm:"birthday"`
}

func (Student) TableName() string {
	return "student"
}
```

### 参考文章
- [配置获取](/docs/20年/11月/01_go项目配置/readme.md)
- [GormV2 API 探究与最佳实践](https://zhuanlan.zhihu.com/p/281594449)

