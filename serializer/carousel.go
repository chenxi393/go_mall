package serializer

import "mail/model"

type Carousel struct {
	ID        uint   `json:"id"`
	ImgPath   string `json:"img_path"`
	ProductId uint   `json:"product_id"`
	CreateAt  int64  `json:"create_at"`
}

func BuildCarousel(carousels *model.Carousel) *Carousel {
	return &Carousel{
		ID:        carousels.ID,
		ImgPath:   carousels.ImgPath,
		ProductId: carousels.ProductID,
		CreateAt:  carousels.CreatedAt.Unix(),
	}
}

func BuildCarousels(carousels []model.Carousel) []Carousel {
	var items []Carousel
	for _, item := range carousels {
		items = append(items, *BuildCarousel(&item))
	}
	return items
}
