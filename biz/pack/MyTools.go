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

func GetUserById(id int64) *api.User {
	var p *api.User = nil
	db, err := gorm.Open(mysql.Open(config.MySQLDSN), &gorm.Config{})
	query.SetDefault(db)
	if err != nil {
		_ = errors.New("db error!")
	} else {
		u, err := query.Q.User.Where(query.User.ID.Eq(id)).Take()
		if err == nil {
			p = User(u)
		}
	}
	return p
}
