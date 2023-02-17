// Code generated by hertz generator.

package feed

import (
	"context"
	"douyin/biz/config"
	"douyin/biz/model/api"
	feed "douyin/biz/model/feed"
	"douyin/biz/model/query"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// Feed .
// @router /douyin/feed/ [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req feed.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(feed.DouyinFeedResponse)

	LastTime := req.GetLatestTime()
	if LastTime == 0 {
		LastTime = time.Now().Unix()
		resp.NextTime = LastTime
	}
	//Token := req.GetToken()

	db, err := gorm.Open(mysql.Open(config.MySQLDSN), &gorm.Config{})
	query.SetDefault(db)

	result, err := query.Q.Video.Where(query.Video.ReleaseTime.Lte(LastTime)).Order(query.Video.ReleaseTime.Desc()).Limit(config.MaxVideoListSize).Find()

	if err != nil {
		resp.StatusMsg = err.Error()
		resp.StatusCode = consts.StatusInternalServerError
		return
	}

	for _, v := range result {
		user, _ := query.Q.User.Where(query.User.ID.Eq(v.UserID)).First()
		LastTime = Min(LastTime, v.ReleaseTime)
		resp.VideoList = append(resp.VideoList, api.Video{
			Id: v.ID,
			Author: &api.User{
				Id:            user.ID,
				Name:          user.Username,
				FollowCount:   1,
				FollowerCount: 1,
				IsFollow:      true,
			},
			PlayUrl:       config.ServerRootUrl + v.PlayURL,
			CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
			FavoriteCount: 1,
			CommentCount:  1,
			IsFavorite:    true,
			Title:         v.Title,
		})
	}

	resp.StatusCode = config.StatusOK
	resp.StatusMsg = consts.StatusMessage(consts.StatusOK)
	c.JSON(consts.StatusOK, resp)
}

func Min(time int64, time2 int64) int64 {
	if time < time2 {
		return time
	}
	return time2
}
