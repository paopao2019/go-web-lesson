package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lesson25/model"
	"log"
	"os"
)


// 全局初始化数据库
func Gorm() *gorm.DB {
	return GormMysql()
}

func GormMysql() *gorm.DB {
	// 数据库mysql
	dsn := "root:p@ss1234@anji@tcp(10.108.26.60:3307)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		os.Exit(0)
		return nil
	} else {
		// 设置参数
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		return db
	}
}


// MysqlTables
//@author: SliverHorn
//@function: MysqlTables
//@description: 注册数据库表专用
//@param: db *gorm.DB

func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.TODO{},
	)
	if err != nil {
		fmt.Printf("初始化表结构失败, err: %v\n", err)
		os.Exit(0)
	}
	log.Println("register table success")
}
