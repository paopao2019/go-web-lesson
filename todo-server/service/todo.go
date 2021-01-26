package service

import (
	"fmt"
	"lesson25/global"
	"lesson25/model"
)

func AddTODO(u model.TODO) (err error){
	err = global.GVA_DB.Create(&u).Error
	return
}

func GetTODOList() (err error, list interface{}, total int64) {
	var todoList []model.TODO
	db := global.GVA_DB.Model(&model.TODO{})
	err = db.Find(&todoList).Error
	db.Count(&total)
	fmt.Printf("List对象: %+v", list)
	return err, todoList, total
}

func UpdateTODO(U *model.TODO, id string) (err error){
	var todo model.TODO
	err = global.GVA_DB.Where("id = ?", id).First(&todo).Updates(map[string]interface{}{
		"Status": U.Status,
		"Title": U.Title,
	}).Error
	return
}

func DeleteTODO(id string) (err error) {
	err = global.GVA_DB.Where("id = ?", id).Delete(&model.TODO{}).Error
	return
}