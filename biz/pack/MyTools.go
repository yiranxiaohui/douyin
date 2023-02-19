package pack

import (
	"crypto/md5"
	"douyin/biz/config"
	"douyin/biz/model/api"
	"douyin/biz/model/query"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetMD5String(s string) (result string) {
	bytes := []byte(s)
	res := md5.Sum(bytes)           //返回值：[Size]byte 数组
	result = fmt.Sprintf("%x", res) //通过fmt.Sprintf()方法格式化数据
	return
}

func IsFollowed(followerId int64, userId int64) (result bool) {
	result = false
	db, err := gorm.Open(mysql.Open(config.MySQLDSN), &gorm.Config{})
	query.SetDefault(db)
	if err != nil {
		_ = errors.New("db error!")
	} else {
		_, err = query.Q.FollowList.Where(query.FollowList.FollowerID.Eq(followerId), query.FollowList.UserID.Eq(userId)).Take()
		result = err == nil
	}
	return
}

func GetUserById(id int64) (*api.User, error) {
	var p *api.User = nil
	var err error = nil
	db, err := gorm.Open(mysql.Open(config.MySQLDSN), &gorm.Config{})
	query.SetDefault(db)
	if err == nil {
		u, err := query.Q.User.Where(query.User.ID.Eq(id)).Take()
		if err == nil {
			p = User(u)
		}
	}
	return p, err
}

type ID struct {
	Id int64
}

func (i ID) GetUserInfoById(id int64) (*api.User, error) {
	var p *api.User = nil
	var err error = nil
	db, err := gorm.Open(mysql.Open(config.MySQLDSN), &gorm.Config{})
	query.SetDefault(db)
	if err == nil {
		u, err := query.Q.User.Where(query.User.ID.Eq(id)).Take()
		if err != nil {
			return nil, err
		}
		p = &api.User{
			Id:            u.ID,
			Name:          u.Username,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      IsFollowed(i.Id, u.ID),
		}
	}
	return p, err
}

func GetLastestMessage(aid int64, bid int64) (*api.Message, error) {
	var p *api.Message = nil
	var err error = nil
	db, err := gorm.Open(mysql.Open(config.MySQLDSN), &gorm.Config{})
	query.SetDefault(db)
	if err == nil {
		msg, err := query.Q.Message.Where(query.Message.FromUserID.Eq(aid), query.Message.ToUserID.Eq(bid)).Or(query.Message.FromUserID.Eq(bid), query.Message.ToUserID.Eq(aid)).Order(query.Message.CreateTime.Desc()).First()
		if err != nil {
			return nil, err
		}
		p = &api.Message{
			Id:         msg.ID,
			ToUserId:   msg.ToUserID,
			FromUserId: msg.FromUserID,
			Content:    msg.Content,
			CreateTime: msg.CreateTime,
		}
	}
	return p, err
}
