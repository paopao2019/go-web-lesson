package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)
// ORM基本认知 github go orm 排名第一的
// ORM = Object Relational Mapping 对象(程序中的对象)关系(真实关系型数据库)映射
/*
结构体 <-> 表
结构体实例 <-> 表中记录行
结构体字段 <-> 表字段
 */

// UserInfo 用户信息
/*
默认和数据库的对应关系 主键、表名、列名的约定
1. GORM 默认会使用名为ID的字段作为表的主键。
2. 表有结构体 名字的小写 驼峰变成_ 并且自动加上s 如 UserInfo 创建的表是 user-infos
3. 字段名字自动是字段名字的小写, 驼峰连接字段变成 _
 */
type UserInfo struct {
	ID uint  // 在DB中的字段是 id 并且是primary key  uint都属于无符号int类型 占8个字节
	Name string
	Gender string
	Hobby string
	Age int            // 在DB中的字段是 age
	TestField string   // 在DB中的字段是 test_field
}

type SysAuthority struct {
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time     `sql:"index"`
	AuthorityId     string         `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	AuthorityName   string         `json:"authorityName" gorm:"comment:角色名"`
	ParentId        string         `json:"parentId" gorm:"comment:父角色ID"`
	DataAuthorityId []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
	Children        []SysAuthority `json:"children" gorm:"-"`
	SysBaseMenus    []SysBaseMenu  `json:"menus" gorm:"many2many:sys_authority_menus;"`
}

type SysBaseMenu struct {
	gorm.Model
	MenuLevel     uint   `json:"-" gorm:"-"`
	ParentId      string `json:"parentId" gorm:"comment:父菜单ID"`
	Path          string `json:"path" gorm:"comment:路由path"`
	Name          string `json:"name" gorm:"comment:路由name"`
	Hidden        bool   `json:"hidden" gorm:"comment:是否在列表隐藏"`
	Component     string `json:"component" gorm:"comment:对应前端文件路径"`
	Sort          int    `json:"sort" gorm:"comment:排序标记"`
	Meta          `json:"meta" gorm:"comment:附加属性"`
	SysAuthoritys []SysAuthority         `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	Children      []SysBaseMenu          `json:"children" gorm:"-"`
	Parameters    []SysBaseMenuParameter `json:"parameters"`
}


type Meta struct {
	KeepAlive   bool   `json:"keepAlive" gorm:"comment:是否缓存"`
	DefaultMenu bool   `json:"defaultMenu" gorm:"comment:是否是基础路由（开发中）"`
	Title       string `json:"title" gorm:"comment:菜单名"`
	Icon        string `json:"icon" gorm:"comment:菜单图标"`
}

type SysBaseMenuParameter struct {
	gorm.Model
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"comment:地址栏携带参数为params还是query"`
	Key           string `json:"key" gorm:"comment:地址栏携带参数的key"`
	Value         string `json:"value" gorm:"comment:地址栏携带参数的值"`
}





//// `User` 属于 `Company`，`CompanyID` 是外键
//type User struct {
//	gorm.Model
//	Name      string
//	CompanyendID int
//	Company   Company `gorm:"foreignKey:CompanyendID"`
//}
//
//type Company struct {
//	ID   int
//	Name string
//}



//// User 有一张 CreditCard，CreditCardID 是外键
//type User struct {
//	gorm.Model
//	CreditCard CreditCard
//}
//
//type CreditCard struct {
//	gorm.Model
//	Number string
//	UserID uint
//}



//// User 有多张 CreditCard，UserID 是外键
//type User struct {
//	gorm.Model
//	CreditCards []CreditCard
//}
//
//type CreditCard struct {
//	gorm.Model
//	Number string
//	UserID uint
//}
//
//type User struct {
//	gorm.Model
//	Languages []*Language `gorm:"many2many:user_languages;"`
//}
//
//type Language struct {
//	gorm.Model
//	Name string
//	Users []*User `gorm:"many2many:user_languages;"`
//}

type User struct {
	gorm.Model
	Username string
	Orders   []Order
}

type Order struct {
	gorm.Model
	UserID uint
	Price  float64
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
	//db.AutoMigrate(&UserInfo{}) // 试过了也可以是UserInfo{}
	//db.Debug().AutoMigrate(&UserInfo{})
	db.Debug().AutoMigrate(&User{})
	db.Debug().AutoMigrate(&Order{})

	// 结构体实例
	//user := User{
	//	Username: "andy",
	//	Orders: []Order {
	//		{Price: 1.10},
	//		{Price: 1.20},
	//	},
	//}
	//db.Debug().Create(&user)


	var u User
	//db.Debug().First(&u)
	db.Debug().Preload("Orders").First(&u)
	fmt.Printf("u values is %+v\n", u)


}
