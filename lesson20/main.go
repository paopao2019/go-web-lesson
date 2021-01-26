package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// gorm的默认约定行为 和 更改约定行为
/*
约定：
1. 约定表名  表名默认就是结构体名称的复数，例如：User 的表名是 users
2. 约定的表字段 是列名由字段名称进行下划线分割(分割驼峰)来生成
3. 主键（Primary Key） GORM 默认会使用名为ID的字段作为表的主键。

string 类型对应 数据库的 varchar size对应varchar的大小
没有设置tag size或者varchar的默认 string是 longtext ?


更改:
1. 表名 不推荐更改 如果要更改可以参考官网
2. 字段名 通过tag column更改 字段名也不推荐更改
3. 主键 通过tag 增加primary_key

----------------------------------------------------------

*/
type User struct {
  gorm.Model   // gorm 自带的一个结构体 里面有四个通用的字段
  Name         string
  Age          sql.NullInt64
  Birthday     *time.Time
  Email        string  `gorm:"type:varchar(100);unique_index;column:email;comment:邮件"`
  Role         string  `gorm:"size:255;comment:角色"` // 设置字段大小为255
  MemberNumber *string `gorm:"unique;not null;comment:会员号"` // 设置会员号（member number）唯一并且不为空
  Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
  Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
  Password    string   `json:"-"  gorm:"comment:用户登录密码"`
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
		db.AutoMigrate(&User{}) // 试过了也可以是UserInfo{}


}

