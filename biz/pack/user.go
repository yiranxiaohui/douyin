package pack

import (
	"douyin/biz/model/api"
	"douyin/biz/model/orm_gen"
)

// Users Convert model.User list to api.User list
func Users(models []*orm_gen.User) []*api.User {
	users := make([]*api.User, 0, len(models))
	for _, m := range models {
		if u := User(m); u != nil {
			users = append(users, u)
		}
	}
	return users
}

// User Convert model.User to api.User
func User(model *orm_gen.User) *api.User {
	if model == nil {
		return nil
	}
	return &api.User{
		Id:            model.ID,
		Name:          model.Username,
		FollowCount:   int64(model.FollowCount),
		FollowerCount: int64(model.FollowerCount),
		IsFollow:      model.IsFollow != 0,
	}
}
