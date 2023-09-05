package service

import (
	"io"
	"mail/config"
	"mime/multipart"
	"os"
	"strconv"
)

func UploadAvatarToLocal(file multipart.File, uid uint, userName string) (string, error) {
	bid := strconv.Itoa(int(uid)) // uint 转换为 string
	basePath := "." + config.My_path.Avatar + "user" + bid + "/"
	if !DirExistOrNot(basePath) {
		CreatDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg" //TODO: 提取file的后缀
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return "", err
	}
	return "user" + bid + "/" + userName + ".jpg", nil
}

func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr) // 获取路径的文件信息
	if err != nil {
		return false
	}
	return s.IsDir() // 判断是不是一个目录
}

func CreatDir(dirName string) bool {
	err := os.MkdirAll(dirName, 0755)
	return err == nil
}
