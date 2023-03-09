package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const dsn = "golang_mall:password@tcp(localhost:3306)/golang_mall?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	fmt.Println("gorm.gen生成代码")
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("创建数据库连接失败: %w", err))
	}
	// 生成实例
	g := gen.NewGenerator(gen.Config{
		ModelPkgPath: "../../server/logic/orm/model",
		OutPath:      "../../server/logic/orm/dal",
		Mode:         gen.WithDefaultQuery | gen.WithoutContext,
	})
	// 设置目标 db
	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
