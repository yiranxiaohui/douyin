// Code generated by hertz generator.

package user

import (
	"context"
	"douyin/biz/config"
	"douyin/biz/model/orm_gen"
	"douyin/biz/model/query"
	"douyin/biz/pack"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	user "douyin/biz/model/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// UserRegister .
// @router /douyin/user/register/ [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.DouyinUserRegisterResponse)

	//取参数
	getUsername := req.GetUsername()
	getPassword := req.GetPassword()

	//设置数据源
	db, err := gorm.Open(mysql.Open(config.MySQLDSN), &gorm.Config{})
	query.SetDefault(db)

	//查询用户名是否已存在
	_, err = query.Q.User.Where(query.User.Username.Eq(getUsername)).Take()

	//不存在同名用户
	if err != nil {
		newUser := orm_gen.User{
			Username:      getUsername,
			Password:      getPassword,
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      0,
		}
		db.Create(&newUser)
		result, err := query.Q.User.Where(query.User.Username.Eq(getUsername)).Take()
		if err != nil {
			resp = &user.DouyinUserRegisterResponse{
				StatusCode: config.StatusInternalServerError,
				StatusMsg:  err.Error(),
			}
		} else {
			userToken := pack.GetMD5String(getUsername + getPassword)
			resp = &user.DouyinUserRegisterResponse{
				StatusCode: config.StatusOK,
				StatusMsg:  consts.StatusMessage(consts.StatusOK),
				UserId:     result.ID,
				Token:      userToken,
			}
		}
	} else {
		resp = &user.DouyinUserRegisterResponse{
			StatusCode: config.StatusInternalServerError,
			StatusMsg:  err.Error(),
		}
	}
	c.JSON(consts.StatusOK, resp)
}
