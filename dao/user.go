package dao

import (
	"context"
	"gorm.io/gorm"
	"mail/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

// 复用连接了的db 优化性能？？ 有点不懂
func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// 根据username  判断表里是不是存在
func (dao *UserDao) ExistByUserName(userName string) (*model.User, bool, error) {
	var user model.User
	// 这里有个问题 如果上述user 声明在返回值里 为结构体指针
	//下面first传递user 会出问题invalid value, should be pointer to struct or slice
	err := dao.DB.Model(&model.User{}).Where("user_name=?", userName).First(&user).Error
	if err != nil {
		//fmt.Println(err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, false, nil
		}
		return &user, true, err
	}
	return &user, true, nil
}

func (dao *UserDao) CreatUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(user).Error
}
