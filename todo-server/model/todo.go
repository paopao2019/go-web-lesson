package model

// 数据库模型
type TODO struct {
	Id int `json:"id"`
	Title string `json:"title" binding:"required"`
	Status bool `json:"status"`
}
