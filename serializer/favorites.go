package serializer

import (
	"context"
	"mail/config"
	"mail/dao"
	"mail/model"
	"mail/pkg/util"
)

type Favorites struct {
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	CreateAt      int64  `json:"create_at"`
	Name          string `json:"name"`
	CategoryID    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	BossID        uint   `json:"boss_id"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
}

func BuildFavorites(favorite *model.Favorite) *Favorites {
	productDao := dao.NewProductDao(context.Background())
	product, err := productDao.GetProductById(favorite.ProductId)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return &Favorites{}
	}
	favorite.Product = *product
	return &Favorites{
		UserID:        favorite.UserId,
		ProductID:     favorite.ProductId,
		CreateAt:      favorite.CreatedAt.Unix(),
		Name:          favorite.Product.Name,
		CategoryID:    favorite.Product.CategoryID,
		Title:         favorite.Product.Title,
		Info:          favorite.Product.Info,
		ImgPath:       config.My_path.Host + config.HttpPort + config.My_path.Product + favorite.Product.ImgPath,
		Price:         favorite.Product.Price,
		DiscountPrice: favorite.Product.DiscountPrice,
		BossID:        favorite.Product.BossID,//应该以商品信息为准
		Num:           favorite.Product.Num,
		OnSale:        favorite.Product.OnSale,
	}
}

func BuildFavoriteses(Favoritess []*model.Favorite) []*Favorites {
	// 不要在循环里对数据库进行操作
	var items []*Favorites
	for _, item := range Favoritess {
		items = append(items, BuildFavorites(item))
	}
	return items
}
