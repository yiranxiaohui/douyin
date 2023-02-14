package main

import (
	"DouyinProject/biz/dal_douyin/mysql_douyin"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./biz/model/query",
		ModelPkgPath: "/orm_gen",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	mysql_douyin.Init()
	g.UseDB(mysql_douyin.DB)

	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
