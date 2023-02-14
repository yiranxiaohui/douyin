package dal_douyin

import (
	"DouyinProject/biz/dal_douyin/mysql_douyin"
	"gorm.io/gen/examples/dal/query"
)

func Init()  {
	mysql_douyin.Init()
	query.SetDefault(mysql_douyin.DB)
}
