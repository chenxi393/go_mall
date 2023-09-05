package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"

	"mail/dao"
)

type Mysql struct {
	//注意viper 解析的结构体字段必须大写
	//这样，Viper 才能够通过反射机制正确解析和赋值。
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
}

type Path struct {
	Host    string `mapstructure:"host"`
	Product string `mapstructure:"product"`
	Avatar  string `mapstructure:"avatar"`
}

type EmailConfig struct {
	ValidEmail string `mapstructure:"valid_email"`
	SMTPHost   string `mapstructure:"smtp_host"`
	SMTPEmail  string `mapstructure:"smtp_email"`
	SMTPPass   string `mapstructure:"smtp_pass"`
}

var (
	AppMode  string
	HttpPort string

	Db_mysql Mysql
	My_path  Path
	Email    EmailConfig

	RedisDb       string
	RedisAddr     string
	RedisPassword string
	RedisDbName   string
)

func Init() {
	viper.SetConfigName("config")    // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config")  // optionally look for config in the working directory
	viper.AddConfigPath("../config") // 取决于main 函数生成的可执行文件在哪
	// 这里的路径后续还得配置 最好用环境变量什么的？？ 不让容易出错 或者决定路径
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	load_config()
	// TODO:读写分离
	// mysql 读 （读写分离） 大部分都是数据库读的操作 主从中的主
	path_read := strings.Join([]string{
		Db_mysql.User,
		":",
		Db_mysql.Password,
		"@tcp(",
		Db_mysql.Host,
		":",
		Db_mysql.Port,
		")/",
		Db_mysql.DBName,
		"?charset=utf8mb4&parseTime=true"}, "")
	// 主从复制
	path_write := strings.Join([]string{
		Db_mysql.User,
		":",
		Db_mysql.Password,
		"@tcp(",
		Db_mysql.Host,
		":",
		Db_mysql.Port,
		")/",
		Db_mysql.DBName,
		"?charset=utf8mb4&parseTime=true"}, "")
	// DSN 格式  data source name
	//username:password@protocol(address)/dbname?param=value
	// 具体命名规则可以看 https://github.com/go-sql-driver/mysql#dsn-data-source-name
	dao.Database(path_read, path_write)
}

func load_config() {
	// 读取配置项
	AppMode = viper.GetString("service.mode")
	HttpPort = viper.GetString("service.http_port")

	err := viper.UnmarshalKey("mysql", &Db_mysql)
	if err != nil {
		panic(fmt.Errorf("fatal error UnmarshalKey: %w", err))
	}

	err = viper.UnmarshalKey("path", &My_path)
	if err != nil {
		panic(fmt.Errorf("fatal error UnmarshalKey: %w", err))
	}

	err = viper.UnmarshalKey("email", &Email)
	if err != nil {
		panic(fmt.Errorf("fatal error UnmarshalKey: %w", err))
	}

	RedisAddr = viper.GetString("redis.host") + viper.GetString("redis.post")
	RedisPassword = viper.GetString("redis.password")
	RedisDbName = viper.GetString("redis.db_name")
}
