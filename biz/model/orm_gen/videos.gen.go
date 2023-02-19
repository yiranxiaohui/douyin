// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package orm_gen

const TableNameVideo = "videos"

// Video mapped from table <videos>
type Video struct {
	ID            int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`    // ID
	UserID        int64  `gorm:"column:user_id;not null" json:"user_id"`               // 作者ID
	PlayURL       string `gorm:"column:play_url;not null" json:"play_url"`             // 播放地址
	CoverURL      string `gorm:"column:cover_url;not null" json:"cover_url"`           // 封面地址
	FavoriteCount int32  `gorm:"column:favorite_count;not null" json:"favorite_count"` // 点赞数
	CommentCount  int32  `gorm:"column:comment_count;not null" json:"comment_count"`   // 评论数
	Title         string `gorm:"column:title;not null" json:"title"`                   // 视频标题
	ReleaseTime   int64  `gorm:"column:release_time;not null" json:"release_time"`     // 提交时间
}

// TableName Video's table name
func (*Video) TableName() string {
	return TableNameVideo
}
