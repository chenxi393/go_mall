package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var _db *gorm.DB //全局的db连接 下面赋值了 和下面的是一样的

func Database(connRead, connWrite string) {
	//自己测试过了 这里gin.Mode 默认就是debug
	// debug模式 会打印详细的日志
	// else那 不打印日志 建表之前试的
	var ormLogger logger.Interface //使用ormLogger来记录ORM操作的日志消息
	if gin.Mode() == "debug" {     //根据代码模式设置不同的日志等级
		ormLogger = logger.Default.LogMode(logger.Info) // Info 会记录详细的日志？
	} else {
		ormLogger = logger.Default
	}
	fmt.Print(connRead)
	// 这里其实还是需要手动创建数据库
	// TODO:调研是不是可以自动创建 虽然可能企业是手动创建
	db, err := gorm.Open(mysql.New(
		mysql.Config{
			DSN:                       connRead,
			DefaultStringSize:         256,
			DisableDatetimePrecision:  true,  //静止Datetime 精度 mysql5.6之前不支持
			DontSupportRenameIndex:    true,  //mysql 5.7之前不支持索引重命名
			DontSupportRenameColumn:   true,  // mysql 8之前的数据库不支持
			SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //单数表
		},
	})
	if err != nil {
		fmt.Println("failed to connect database")
	}
	sqlDB, _ := db.DB()        //获取底层连接对象
	sqlDB.SetMaxIdleConns(10)  // 最大空闲（idle）连接数
	sqlDB.SetMaxOpenConns(100) //设置最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db

	// 主从配置 https://gorm.io/zh_CN/docs/dbresolver.html
	// gorm 文档建议写操作使用源库sources（也就是主） 读使用副本库 replicas（从）
	//读写分离代理
	_ = _db.Use(dbresolver.Register(
		dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(connWrite)},
			Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)},
			Policy:   dbresolver.RandomPolicy{}, //负载均衡的策略
		}))
	//有migration()函数时，数据库表格的创建和更新是自动进行的，减少了手动管理的工作量。
	// 一般来说 企业有dbms的 数据库已经创建好了 不需要这个
	migration()
}

// 这一段还不知道干嘛的
func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
