package dao

import (
	"fmt"
	"mail/model"
)

// 这里其实就是自动建表
// 总结：进行数据库迁移操作。具体的行为取决于连接的数据库的状态和模型定义的结构。
// 如果数据库中不存在相关表格，将会创建对应的表格。
// 如果已经存在相关表格，将会根据需要进行更新操作。
// 无论数据库是否有数据，迁移操作不会删除或修改已有的数据。
func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.Address{},
		&model.Admin{},
		&model.BasePage{},
		&model.Carousel{},
		&model.Cart{},
		&model.Category{},
		&model.Product{},
		&model.ProductImg{},
		&model.User{},
		&model.Favorite{},
		&model.Notice{},
		&model.Order{},
	)
	if err != nil {
		fmt.Println("err", err)
	}
}
