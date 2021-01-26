package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lesson25/model"
	"lesson25/model/request"
	"lesson25/model/response"
	"lesson25/service"
	"strconv"
)

func TODOGet(c *gin.Context) {
	// 响应的数据应该是个list
	if err, list, total := service.GetTODOList(); err != nil {
		//c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
			Total: total,
			Page: 1,
			PageSize: 10,
		}, "获取成功", c)

	}
}

// 增加待办项
func TODOPost(c *gin.Context) {
	var todo request.AddTODO
	// 1. 接收请求参数
	if err :=c.ShouldBindJSON(&todo); err != nil {
		fmt.Printf("todo Post 方法, 参数有误, error: %v", err)
		response.FailWithMessage("参数绑定失败", c)
		return
	}
	//U := &model.SysUser{Username: user.Username, Password: user.Password}
	U := &model.TODO{Title: todo.Title, Status: todo.Status}

	// 2. 插入数据库  应该抽离层业务层
	err := service.AddTODO(*U)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(todo, c)
	}

}

// 更新状态
func TODOPut(c *gin.Context) {

	// 1. 接受请求参数
	id := c.Param("id")

	var todo request.AddTODO
	if err :=c.ShouldBindJSON(&todo); err != nil {
		fmt.Printf("todo PUT 方法, 参数有误, error: %v", err)
		response.FailWithMessage("参数绑定失败", c)
		return
	}
	// 2. 业务逻辑处理
	U := &model.TODO{Title: todo.Title, Status: todo.Status}
	U.Id , _ = strconv.Atoi(id)

	// 3. 响应
	if err := service.UpdateTODO(U, id); err != nil {
		//c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(*U, c)

	}

}


// 删除TODO LIST
func TODODelete(c *gin.Context) {
	id := c.Param("id")
	err := service.DeleteTODO(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}else{
		response.OkWithMessage("删除成功", c)
	}

}

