// Code generated by hertz generator.

package favorite

import (
	"context"

	favorite "douyin/biz/model/favorite"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// FavoriteList .
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite.DouyinFavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(favorite.DouyinFavoriteListResponse)

	c.JSON(consts.StatusOK, resp)
}