// Code generated by hertz generator.

package publish

import (
	"context"
	"douyin/biz/config"
	"douyin/biz/model/orm_gen"
	"douyin/biz/model/query"
	"douyin/biz/pack"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"os"
	"strconv"
	"time"

	publish "douyin/biz/model/publish"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// PublishAction .
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req publish.DouyinPublishActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(publish.DouyinPublishActionResponse)
	defer c.JSON(consts.StatusOK, resp)

	getData := req.GetData()
	getToken := req.GetToken()
	getTitle := req.GetTitle()

	userdata, err := pack.ParseToken(getToken)

	if err != nil {
		resp.StatusMsg = err.Error()
		resp.StatusCode = consts.StatusInternalServerError
		return
	}

	nowUnix := time.Now().Unix()

	VideoFilePath := config.ServerVideoPath + strconv.FormatInt(userdata.ID, 10) + "_" + strconv.FormatInt(nowUnix, 10) + ".mp4"

	out, err := os.Create(VideoFilePath)

	if err != nil {
		resp.StatusMsg = err.Error()
		resp.StatusCode = consts.StatusInternalServerError
		return
	}

	data, err := getData.Open()

	if err != nil {
		resp.StatusMsg = err.Error()
		resp.StatusCode = consts.StatusInternalServerError
		return
	}

	_, err = io.Copy(out, data)

	if err != nil {
		resp.StatusMsg = err.Error()
		resp.StatusCode = consts.StatusInternalServerError
		return
	}

	db, err := gorm.Open(mysql.Open(config.MySQLDSN), &gorm.Config{})
	query.SetDefault(db)

	newVideo := &orm_gen.Video{
		ID:           pack.GetSnowflakeId(),
		UserID:       userdata.ID,
		PlayURL:      VideoFilePath,
		CoverURL:     "",
		CommentCount: 0,
		Title:        getTitle,
		ReleaseTime:  nowUnix,
	}

	db.Create(newVideo)

	resp.StatusCode = config.StatusOK
	resp.StatusMsg = consts.StatusMessage(consts.StatusOK)

}
