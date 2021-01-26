package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


// Mysql　查询
/*
1. 简单查询first 赋值 查询所有记录
2. where条件查询  ? 是占位符
3. find查询 里面的conditions ?是条件匹配占位符
4. 按字段查询 select
4. 高级查询

实际如何使用
1. 每次会加.Error  Error类型
2. Find中会加条件，可能不使用where做链式查找

3. join 等操作 内联操作 等用到的时候才来看文档
 */

type User struct {
	Id int
	Name string
	//Age int `gorm:"default:18"`
	//Age *int  `gorm:"default:18"`
	Age sql.NullInt64 `gorm:"default:18"` // sql.NullString 实现了Scanner/Valuer接口
}

func main() {
	// 数据库mysql
	dsn := "root:p@ss1234@anji@tcp(10.108.26.60:3307)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Mysql数据库无法连接, error: %v", err)
		return
	} else {
		// 设置参数
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
	}

	// 自动迁移 会自动和数据库对应 创建数据库表  如果表存在会自动重新更新数据库如果Model有更改
	db.AutoMigrate(&User{})

	// 创建记录
	//u1 := User{Name: "andy", Age: sql.NullInt64{Int64: 10, Valid: true}}
	//u2 := User{Name: "andy", Age: sql.NullInt64{Int64: 20, Valid: true}}
	//db.Create(&u1)
	//db.Create(&u2)

	// 查询
	var u User
	// 查询第一条记录 order by primary key(id)
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	db.Debug().First(&u)
	fmt.Printf("user:%#v\n", u.Age.Int64)

	// 查询所有记录
	var userList []User
	db.Find(&userList)
	fmt.Printf("userList:%+v\n", userList) //[{Id:1 Name:andy Age:{Int64:10 Valid:true}} {Id:2 Name:andy Age:{Int64:20 Valid:true}}]

	// where查询
	var findUserList []User
	//db.Where("age = ?", 10).First(&u)
	// 字段错误会有报错
	err = db.Where("age = ?", 10).Debug().Find(&findUserList).Error
	if err != nil {
		fmt.Printf("%T, %v", err, err)
		return
	}
	fmt.Printf("findUserList:%+v\n", findUserList)

	// match counter
	userCounter := new(int64)
	// 查询大于等于10的总数
	//db.Where("age >= ?", 10).Find(&findUserList).Count(userCounter)
	db.Find(&findUserList, "age >= ?", 10).Count(userCounter)
	fmt.Printf("userCounter:%v\n", *userCounter)  // 2

	// select 字段
	var selectUserList []User
	db.Where("age >= ?", 10).Select("name").Debug().Find(&selectUserList)
	// 输出结果 映射到 结构体上结构体其他字段会有结构体的默认值
	fmt.Printf("selectUserList:%+v\n", selectUserList)

	// limit 输出
	var limitUserList []User
	db.Limit(3).Find(&limitUserList)
	fmt.Printf("limitUserList: %+v\n", limitUserList)


}