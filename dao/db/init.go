/**
 @author: RedCrazyGhost
 @date: 2022/5/31

**/

package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

// Init 数据库连接初始化
func Init() {
	var err error
	dsn := "root:123456@tcp(81.69.173.253:3306)/pangolin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(&Douban{})
	if err != nil {
		return
	}
	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}
}
