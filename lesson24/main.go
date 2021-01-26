package main

// 删除
/*
1. gorm是软删除
2. 	如果一个 model 有 DeletedAt 字段，他将自动获得软删除的功能！
当调用 Delete 方法时， 记录不会真正的从数据库中被删除， 只会将DeletedAt 字段的值会被设置为当前时间

3. 没有嵌入默认的 gorm.Model 就没有软删除的概念了
1. 根据主键去删除
2. 根据where条件去删除

 */
import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

	// 安主键删除一条指定的数据
	var user User
	user.Id = 1
	// UPDATE `users` SET `deleted_at`='2021-01-20 14:21:11.916' WHERE `users`.`id` = 1 AND `users`.`deleted_at` IS NULL
	db.Debug().Delete(&user)

	// 软删除操作
	//UPDATE `users` SET `deleted_at`='2021-01-20 14:16:24.268' WHERE age >= 10 AND `users`.`deleted_at` IS NULL
	db.Debug().Where("age >= ?", 10).Delete(&User{})

	// 物理删除
	//  DELETE FROM `users` WHERE age >= 10
	db.Debug().Unscoped().Where("age >= ?", 10).Delete(&User{})

}
