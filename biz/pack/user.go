package pack

import (
	"douyin/biz/model/api"
	"douyin/biz/model/orm_gen"
)

func User(user *orm_gen.User) *api.User {
	return &api.User{
		Id:            user.ID,
		Name:          user.Username,
		FollowCount:   int64(user.FollowCount),
		FollowerCount: int64(user.FollowerCount),
		IsFollow:      user.IsFollow != 0,
	}
}
