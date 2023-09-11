package dao

import (
	"context"
	"mail/model"

	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{newDBClient(ctx)}
}

func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

func (dao *NoticeDao) GetNoticeById(id uint) (*model.Notice, error) {
	var notice model.Notice
	err := dao.DB.Model(&model.Notice{}).Where("id = ?", id).First(&notice).Error
	return &notice, err
}
