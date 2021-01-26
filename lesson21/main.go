package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// gorm创建

/*
掌握:
1. 默认值
2. 在结构体中 定义的字段如果不传值会有默认值如 int 是0 string是"",那怎么样映射到DB中呢？
	是映射"" 还是Null呢？ 这些不一样的需求都怎么满足呢？
3. 可以调试 生成的语句？

结论:
1. 默认行为 将所有的结构体默认值插入到数据中， 这是结构体特性，所以会将映射到数据字段上
2. 通过过tag定义字段的默认值，在创建记录时候生成的 SQL 语句会排除没有值或值为 零值 的字段。
	在将记录插入到数据库后，Gorm会从数据库加载那些字段的默认值。
	过tag定义字段的默认值，在创建记录时候生成的 SQL 语句会排除没有值或值为 零值 的字段。
	在将记录插入到数据库后，Gorm会从数据库加载那些字段的默认值
3. 我实际希望达到的需求是 我传了字段就按字段的来, 如Age:0 应该是传了值，而不是去用default默认值，
	当我没有传递字段的时候，才使用gorm的默认值。
	实现方式1： 字段使用指针
	实现方式2： 使用Scanner/Valuer接口方式实现零值存入数据库

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

	//var age *int   // 这是错误的指针赋值方式 指针一定要new分配内存空间 才可以赋值，所以一定要new
	//*age = 11
	age := new(int)
	*age = 10
	//u1 := User{Name: "张三", Age: age}
	u1 := User{Name: "张三", Age: sql.NullInt64{Int64: 0,Valid: true}} // 该条记录age字段的值就是''
	//u2 := User{Name: "张三", Age: new(int)}
	u2 := User{Name: "张三"}


	// 创建记录
	db.Create(&u1)
	db.Create(&u2)

	//db.Debug().Create(&u2)  // 调试 看看orm转换成的sql语句 可能会有Error Duplicate 提示，因为也是正在执行


}
