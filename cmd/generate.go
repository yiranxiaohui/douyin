package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const MySQLDSN = "root:qazwsx147258@(localhost:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"

func main() {

	// 连接数据库
	db, err := gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	// 生成实例
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./biz/model/query",
		ModelPkgPath: "/orm_gen",
		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db)

	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
