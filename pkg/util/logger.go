package util

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

var LogrusObj *logrus.Logger

func init() {
	src, _ := setOutPutFile()
	if LogrusObj != nil {
		LogrusObj.SetOutput(src)
		return
	}
	logger := logrus.New()
	logger.SetOutput(src)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 后续ELK体系  添加
	// logger.AddHook()
	LogrusObj = logger
}

// 希望日志能按天区分
func setOutPutFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	// 获取当前目录的绝对路径
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	// 若目录不存在 则创建
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	//直接用下面的指令 不存在自动创建
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	// 更改文件权限
	if err != nil {
		return nil, err
	}
	err = os.Chmod(fileName, 0777)
	if err != nil {
		log.Fatal(err)
	}
	return src, nil
}
