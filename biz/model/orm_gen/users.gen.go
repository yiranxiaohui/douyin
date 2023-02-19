// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package orm_gen

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID            int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`    // ID
	Username      string `gorm:"column:username;not null" json:"username"`             // 用户名
	Password      string `gorm:"column:password;not null" json:"password"`             // 密码
	FollowCount   int64  `gorm:"column:follow_count;not null" json:"follow_count"`     // 关注总数
	FollowerCount int64  `gorm:"column:follower_count;not null" json:"follower_count"` // 粉丝总数
	IsFollow      int32  `gorm:"column:is_follow;not null" json:"is_follow"`           // 是否关注
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
