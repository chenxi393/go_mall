package dao

import (
	"context"
	"gorm.io/gorm"
	"mail/model"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{newDBClient(ctx)}
}

func (dao *ProductDao) CreatProduct(Product *model.Product) error {
	// 这里和前面create 里的已经是指针类型 还需要加指针吗 GPT认为不需要 视频里加了
	// 遇到很多次上面类似的疑问了 返回一组又怎么办 感觉还是得看文档 找个时间仔细看文档
	return dao.DB.Model(&model.Product{}).Create(Product).Error
}

func (dao *ProductDao) GetProductById(id uint) (*model.Product, error) {
	var Product model.Product
	err := dao.DB.Model(&model.Product{}).Where("id = ?", id).First(&Product).Error
	return &Product, err
}

func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (int64, error) {
	var cnt int64
	// 也可以传 除了where ? int
	// 也可以map 键值对 或者struct 具体可以见文档
	// https://gorm.io/zh_CN/docs/query.html#Struct-amp-Map-%E6%9D%A1%E4%BB%B6
	err := dao.DB.Model(&model.Product{}).Where(condition).Count(&cnt).Error
	return cnt, err
}

func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, page model.BasePage) ([]model.Product, error) {
	var products []model.Product
	// 偏移量应该是当前页数-1 * 一页的数量
	err := dao.DB.Model(&model.Product{}).Where(condition).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&products).Error
	return products, err
}

func (dao *ProductDao) SearchProduct(info string, page model.BasePage) ([]model.Product, error) {
	var products []model.Product
	err := dao.DB.Model(&model.Product{}).Where("title LIKE ? OR info LIKE ? ", "%"+info+"%", "%"+info+"%").Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&products).Error
	return products, err
}
