package serializer

import (
	"mail/config"
	"mail/model"
)

type ProductImag struct {
	ProductId uint   `json:"product_id"`
	ImagPath  string `json:"img_path"`
}

func BuildProductImag(item *model.ProductImg) ProductImag {
	return ProductImag{
		ProductId: item.ProductID,
		ImagPath:  config.My_path.Host + config.HttpPort + config.My_path.Product + item.ImgPath,
	}
}

func BuildImags(items []*model.ProductImg) []ProductImag {
	imags := make([]ProductImag, 0, len(items))
	for _, item := range items {
		imags = append(imags, BuildProductImag(item))
	}
	return imags

}
