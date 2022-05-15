package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:LLQtT$3v@tcp(127.0.0.1:3306)/homework?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	g := gen.NewGenerator(gen.Config{OutPath: "../dal/query"})

	g.UseDB(db)
	g.ApplyBasic(
		g.GenerateModelAs("user", "User",
			gen.FieldType("uuid", "string"),
			gen.FieldType("name", "string"),
			gen.FieldType("age", "int64"),
			gen.FieldType("version", "int64"),
		))

	// apply diy interfaces on structs or table models
	// 如果想给某些表或者model生成自定义方法，可以用ApplyInterface，第一个参数是方法接口，可以参考DIY部分文档定义
	//g.ApplyInterface(func(method model.Method) {}, model.User{}, g.GenerateModel("company"))

	// execute the action of code generation
	g.Execute()
}
