// Code generated by hertz generator.

package comment

import (
	"context"
	"douyin/biz/config"
	"douyin/biz/model/api"
	"douyin/biz/model/query"
	"douyin/biz/pack"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"

	comment "douyin/biz/model/comment"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CommentList .
// @router /douyin/comment/list [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req comment.DouyinCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(comment.DouyinCommentListResponse)
	defer c.JSON(consts.StatusOK, resp)

	db, err := gorm.Open(mysql.Open(config.MySQLDSN), &gorm.Config{})
	query.SetDefault(db)

	commentIdList, err := query.Q.CommentList.Select(query.CommentList.ID).Where(query.CommentList.VideoID.Eq(req.GetVideoId())).Find()
	if err != nil {
		resp.StatusMsg = err.Error()
		resp.StatusCode = consts.StatusInternalServerError
		return
	}
	for _, v := range commentIdList {
		com, err := query.Q.Comment.Where(query.Comment.ID.Eq(v.ID)).Take()
		if err != nil {
			resp.StatusMsg = err.Error()
			resp.StatusCode = consts.StatusInternalServerError
			return
		}
		user, err := pack.GetUserById(com.UserID)
		if err != nil {
			resp.StatusMsg = err.Error()
			resp.StatusCode = consts.StatusInternalServerError
			return
		}
		resp.CommentList = append(resp.CommentList, api.Comment{
			Id:         com.ID,
			User:       user,
			Content:    com.Text,
			CreateDate: time.Unix(com.CreateDate, 0).Format("2006-01-02 15:04:05"),
		})
	}
	resp.StatusCode = config.StatusOK
	resp.StatusMsg = consts.StatusMessage(consts.StatusOK)
}
