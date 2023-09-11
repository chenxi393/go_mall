package service

import (
	"context"
	"mail/dao"
	"mail/model"
	"mail/pkg/e"
	"mail/pkg/util"
	"mail/serializer"
	"mime/multipart"
	"strconv"
)

type ProductService struct {
	// 感觉这里面ID 肯定是没有东西的
	// ID是鉴权获取的
	Id             uint   `form:"id" json:"id"`
	Name           string `form:"name" json:"name"`
	CategoryId     uint   `form:"category_id" json:"category_id"`
	Title          string `form:"title" json:"title"`
	Information    string `form:"info" json:"info"`
	ImgPath        string `form:"img_path" json:"img_path"`
	Price          string `form:"price" json:"price"`
	Discount_price string `form:"discount_price" json:"discount_price"`
	OnSale         bool   `json:"on_sale" form:"on_sale"`
	Num            int    `json:"num" form:"num"`
	model.BasePage        //自己实验过了 这种写法不会序列化时不会多出一层结构体嵌套
}

func (service *ProductService) Create(ctx context.Context, uId uint, files []*multipart.FileHeader) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	boss, _ := userDao.GetUserById(uId)
	// 以第一张作为封面图
	tmp, _ := files[0].Open() // 这里要是用户没上传图片会越界 TODO:fix
	path, err := UploadProductToLocal(tmp, uId, service.Name)
	if err != nil {
		code = e.ErrorProductImagUploadError
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	product := &model.Product{
		Name:          service.Name,
		CategoryID:    service.CategoryId,
		Title:         service.Title,
		Info:          service.Information,
		ImgPath:       path,
		Price:         service.Price,
		DiscountPrice: service.Discount_price,
		OnSale:        true,
		Num:           service.Num,
		BossID:        uId,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreatProduct(product)
	if err != nil {
		code = e.ErrorProductImagUploadError
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	for index, file := range files {
		num := strconv.Itoa(index)
		tmp, _ := file.Open()
		productImagDao := dao.NewProductImgDaoByDB(productDao.DB)
		path, err = UploadProductToLocal(tmp, uId, service.Name+num)
		if err != nil {
			code = e.ErrorProductImagUploadError
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		productImg := model.ProductImg{
			ProductID: product.ID,
			ImgPath:   path,
		}
		err = productImagDao.Create(&productImg)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}

func (service *ProductService) List(ctx context.Context) serializer.Response {
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	productDao := dao.NewProductDao(ctx)
	// 像视频用map有个好处就是可以复用 下次传个map而不是再写一个函数
	condition := make(map[string]interface{})
	if service.CategoryId != 0 {
		condition["category_id"] = service.CategoryId
	}
	total, err := productDao.CountProductByCondition(condition)
	if err != nil {
		util.LogrusObj.Infoln(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 我觉得它的意思应该是拿到total 然后goroutine去做分页展示
	// 后面这段 视频用了goroutine 然后wait 我觉得有点脱裤子放屁 没什么用
	products, err := productDao.ListProductByCondition(condition, service.BasePage)
	if err != nil {
		util.LogrusObj.Infoln(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}

// 一般生产中使用 ElasticSearch 进行搜索
func (service *ProductService) Search(ctx context.Context) serializer.Response {
	code := e.Success
	productDao := dao.NewProductDao(ctx)
	// 这里的意思应该是模糊搜索
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	products, err := productDao.SearchProduct(service.Information, service.BasePage)
	if err != nil {
		util.LogrusObj.Infoln(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(len(products)))
}

func (service *ProductService) Get(ctx context.Context, id string) serializer.Response {
	code := e.Success
	pid, _ := strconv.Atoi(id)
	daoProduct := dao.NewProductDao(ctx)
	product, err := daoProduct.GetProductById(uint(pid))
	if err != nil {
		util.LogrusObj.Infoln(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	product.AddView()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}
