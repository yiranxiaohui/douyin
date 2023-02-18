package pack

import (
	"douyin/biz/model/api"
	"douyin/biz/model/orm_gen"
	"douyin/biz/model/query"
)

func Videos(models []*orm_gen.Video)  []*api.Video{
	videos := make([]*api.Video,0, len(models))
	for _, m := range models {
		println("models......")
		if u := Video(m); u != nil {
			println("Video......")
			videos = append(videos, u)
		}
	}
	return videos
}

func Video(model *orm_gen.Video)  *api.Video{
	if model == nil {
		return nil
	}
	first, err := query.User.Where(query.User.ID.In(model.UserID)).First()
	if err != nil {
		return nil
	}
	return &api.Video{
		Id:           model.ID,
		Author:       User(first),
		PlayUrl:      model.PlayURL,
		CoverUrl:     model.CoverURL,
		CommentCount: int64(model.CommentCount),
		Title: model.Title,
	}
}