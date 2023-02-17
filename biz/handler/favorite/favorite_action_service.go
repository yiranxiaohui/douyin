// Code generated by hertz generator.

package favorite

import (
	"context"
	"douyin/biz/config"
	"douyin/biz/model/orm_gen"
	"douyin/biz/model/query"
	"douyin/biz/pack"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	favorite "douyin/biz/model/favorite"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// FavoriteAction .
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite.DouyinFavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(favorite.DouyinFavoriteActionResponse)

	//设置数据库
	db, err := gorm.Open(mysql.Open(config.MySQLDSN), &gorm.Config{})
	query.SetDefault(db)

	//获取前端数据
	getToken := req.Token
	getVideoId := req.VideoId
	getActionType := req.ActionType
	isEorror := 0 //标记是否正确

	//通过ActionType确定是点赞还是取消
	if getActionType == 1 { //点赞
		myClaims, err := pack.ParseToken(getToken)
		userFavor, err := query.User.Where(query.User.ID.Eq(myClaims.ID)).First()
		if err != nil {
			isEorror = 1
		} else {
			//查找是否有重复数据
			first, err := query.Favorite.Where(query.Favorite.UserID.Eq(myClaims.ID).AddCol(query.Favorite.VideoID.Eq(getVideoId))).First()
			if err != nil || first == nil {
				query.Favorite.Create(&orm_gen.Favorite{UserID: userFavor.ID, VideoID: getVideoId})
				isEorror = 0
			} else {
				isEorror = 1
			}
		}
	} else { //取消
		myClaims, err := pack.ParseToken(getToken)
		userDisFavor, err := query.User.Where(query.User.ID.Eq(myClaims.ID)).First()
		if err != nil {
			isEorror = 0
		} else {
			query.Favorite.Unscoped().Delete(&orm_gen.Favorite{UserID: userDisFavor.ID, VideoID: getVideoId})
			isEorror = 1
		}
	}

	if isEorror == 1 {
		resp = &favorite.DouyinFavoriteActionResponse{
			StatusCode: config.StatusInternalServerError,
			StatusMsg:  consts.StatusMessage(consts.StatusInternalServerError),
		}
	} else {
		resp = &favorite.DouyinFavoriteActionResponse{
			StatusCode: config.StatusOK,
			StatusMsg:  consts.StatusMessage(consts.StatusOK),
		}
	}

	println(resp.StatusCode)
	c.JSON(consts.StatusOK, resp)
}
