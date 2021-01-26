package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// update 更新数据库
/*
1. 简单用法 返回值 和 error 捕捉 Save()  默认会更新该对象的所有字段，即使你没有赋值。
2. 如果你只希望更新指定字段，可以使用Update或者Updates
3. 默认会触发钩子函数 ，更新 UpdatedAt 时间戳

实际用是
1. 先查找 在修改 在save或者update
2. where 批量更新

就是要满足 你常用的update语句

理解
db.Model  就是你指定要更新的记录
 // update all users's name to `hello`
   db.Model(&User{}).Update("name", "hello")
   // if user's primary key is non-blank, will use it as condition, then will only update the user's name to `hello`
   db.Model(&user).Update("name", "hello")
*/
type User struct {
	gorm.Model   // 继承公共的字段
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

	// 更新操作
	var user User
	//user.Name = "coco"
	//// 新建 会直接save一条新记录
	//err = db.Save(&user).Error
	//if err != nil {
	//	fmt.Printf("save err, %v", err)
	//	return
	//}

	// 正常应该是 查找后 在save
	db.First(&user)
	user.Age = sql.NullInt64{Int64: 100, Valid: true}
	user.Name = "cherry"
	// 会更新所有字段
	db.Save(&user)
	// 更新指定字段
	// UPDATE `users` SET `name`='ck',`updated_at`='2021-01-20 13:32:40.266' WHERE `id` = 1
	db.Model(&user).Debug().Update("name","ck")


	// age >= 20 Name 全部更新为jack
	// UPDATE `users` SET `name`='jack',`updated_at`='2021-01-20 13:45:05.179' WHERE age >= 20
	db.Model(User{}).Debug().Where("age >= ?", 20).Update("name", "jack")

	// 使用 map 更新多个属性
	//updateInfo := make(map[string]interface{})
	//updateInfo["name"] = "andy"
	//updateInfo["age"] = sql.NullInt64{Int64: 1, Valid: true}
	updateInfo := map[string]interface{} {
		"name": "lucy",
		"age": sql.NullInt64{Int64: 2, Valid: true},
	}
	db.Model(User{}).Debug().Where("age >= ?", 20).Updates(updateInfo)

}
