package serializer

import "mail/model"

type Category struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreateAt     int64  `json:"create_at"`
}

func BuildCategory(categorys *model.Category) *Category {
	return &Category{
		ID:           categorys.ID,
		CategoryName: categorys.CategoryName,
		CreateAt:     categorys.CreatedAt.Unix(),
	}
}

func BuildCategorys(Categorys []*model.Category) []*Category {
	var items []*Category
	for _, item := range Categorys {
		items = append(items, BuildCategory(item))
	}
	return items
}
