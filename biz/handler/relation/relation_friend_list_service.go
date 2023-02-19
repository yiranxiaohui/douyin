// Code generated by hertz generator.

package relation

import (
	"context"
	"douyin/biz/config"
	"douyin/biz/model/api"
	"douyin/biz/model/query"
	"douyin/biz/pack"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	relation "douyin/biz/model/relation"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// RelationFriendList .
// @router /douyin/relation/friend/list [GET]
func RelationFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.DouyinRelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(relation.DouyinRelationFriendListResponse)
	defer c.JSON(consts.StatusOK, resp)

	db, err := gorm.Open(mysql.Open(config.MySQLDSN), &gorm.Config{})
	query.SetDefault(db)

	getToken := req.GetToken()
	claims, err := pack.ParseToken(getToken)
	if err != nil {
		resp.StatusMsg = err.Error()
		resp.StatusCode = config.StatusInternalServerError
		return
	}
	UserId := pack.ID{claims.ID}

	followLists, err := query.Q.FollowList.Where(query.FollowList.FollowerID.Eq(UserId.Id)).Find()
	if err != nil {
		resp.StatusMsg = err.Error()
		resp.StatusCode = config.StatusInternalServerError
		return
	}
	followerLists, err := query.Q.FollowList.Where(query.FollowList.UserID.Eq(UserId.Id)).Find()
	if err != nil {
		resp.StatusMsg = err.Error()
		resp.StatusCode = config.StatusInternalServerError
		return
	}

	m := make(map[int64]bool)

	for _, v := range followLists {
		m[v.UserID] = true
	}

	for _, v := range followerLists {
		if m[v.FollowerID] {
			friend, err := UserId.GetUserInfoById(v.FollowerID)
			if err != nil {
				resp.StatusMsg = err.Error()
				resp.StatusCode = config.StatusInternalServerError
				return
			}
			resp.UserList = append(resp.UserList, api.FriendUser{
				Id:            friend.Id,
				Name:          friend.Name,
				FollowCount:   friend.FollowCount,
				FollowerCount: friend.FollowerCount,
				IsFollow:      friend.IsFollow,
				//待实现
				Message: "这是一条测试消息",
				//0为请求，1为发送
				MsgType: 0,
			})
		}
	}
	resp.StatusCode = config.StatusOK
	resp.StatusMsg = consts.StatusMessage(consts.StatusOK)
}
