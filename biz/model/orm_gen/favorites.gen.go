// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package orm_gen

const TableNameFavorite = "favorites"

// Favorite mapped from table <favorites>
type Favorite struct {
	VideoID int64 `gorm:"column:video_id;not null" json:"video_id"` // 喜欢的视频ID
	UserID  int64 `gorm:"column:user_id;not null" json:"user_id"`   // 用户ID
}

// TableName Favorite's table name
func (*Favorite) TableName() string {
	return TableNameFavorite
}
