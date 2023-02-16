// Code generated by hertz generator.

package user

import (
	"context"
	"douyin/biz/config"
	"douyin/biz/model/query"
	"douyin/biz/pack"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"

	user "douyin/biz/model/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// User .
// @router /douyin/user/ [GET]
func User(ctx context.Context, c *app.RequestContext) {

	var err error
	var req user.DouyinUserRequest
	err = c.BindAndValidate(&req)

	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.DouyinUserResponse)

	getUserId := req.GetUserId()
	getToken := req.GetToken()

	db, err := gorm.Open(mysql.Open(config.MySQLDSN), &gorm.Config{})

	query.SetDefault(db)
	result, err := query.Q.User.Where(query.User.ID.In(getUserId)).First()
	if err != nil {
		_ = fmt.Errorf("用户ID出错！%v", err)
	}

	token, _ := pack.MakeToken(result.ID, result.Password)
	if strings.Compare(getToken, token) == 0 {
		resp = &user.DouyinUserResponse{User: pack.User(result), StatusMsg: consts.StatusMessage(consts.StatusOK), StatusCode: config.StatusOK}
	} else {
		resp = &user.DouyinUserResponse{User: nil, StatusCode: config.StatusInternalServerError, StatusMsg: "用户ID出错！"}
	}

	c.JSON(consts.StatusOK, resp)
}
