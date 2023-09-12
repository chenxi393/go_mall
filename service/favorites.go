package service

import (
	"context"
	"mail/dao"
	"mail/model"
	"mail/pkg/e"
	"mail/pkg/util"
	"mail/serializer"
)

type FavoriteService struct {
	ProductId      uint `form:"product_id" json:"product_id"`
	BossId         uint `form:"boss_id" json:"boss_id"`
	FavoriteId     uint `form:"favorite_id" json:"favorite_id"`
	model.BasePage      //这里进行一个分页的操作 TODO
}

// 注意create 之前要判断存不存在 这个address也没有判断（地址还好） 收藏夹一定要判断
// 还有个问题是 其实应该让用户传boss_id 商品就是自带商家的 
func (service *FavoriteService) Create(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	favorite := model.Favorite{
		UserId:    uid,
		ProductId: service.ProductId,
		BossId:    service.BossId,
	}
	favoritesDao := dao.NewFavoritesDao(ctx)
	exist, err := favoritesDao.IsExist(service.ProductId, uid)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	} else if exist {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    "收藏夹已存在该商品",
		}
	}
	// 如果像视频那样 在创建之前判断 商品和创建该商品的用户有没有存在 那就不需要外键了
	// 这里TODO吧 暂时不像视频里那样写 后续再优化
	err = favoritesDao.Create(&favorite)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *FavoriteService) Delete(ctx context.Context, uid uint, fid string) serializer.Response {
	code := e.Success
	favoritesDao := dao.NewFavoritesDao(ctx)
	err := favoritesDao.DeleteById(uid, fid)
	if err != nil {
		print(32131)
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *FavoriteService) GetAll(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	favoritesDao := dao.NewFavoritesDao(ctx)
	favorites, err := favoritesDao.Getfavorites(uid)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavoriteses(favorites), uint(len(favorites)))
}
