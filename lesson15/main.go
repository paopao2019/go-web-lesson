package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)


// 文件上传
// 1. 实现单个文件上传 -> 保存到本地
// 2. 实现多个文件上传 -> 保存到本地

func main() {
	r := gin.Default()
	// 解析html
	r.LoadHTMLFiles("upload.html", "mulupload.html")

	// 单个文件上传的页面
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})

	// 多个文件上传
	r.GET("/mulupload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "mulupload.html", nil)
	})

	// 单个文件上传
	r.POST("/upload", func(c *gin.Context) {
		f1, err := c.FormFile("f1")  // *FileHeader // 从请求中携带的参数要一样的
		if err != nil {
			fmt.Printf("获取文件失败, error: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			dst := path.Join("./", f1.Filename)
			c.SaveUploadedFile(f1, dst)
		}

		c.String(http.StatusOK, "上传文件成功 " + f1.Filename)

	})
	// 多个文件上传
	r.POST("/mulupload", func(c *gin.Context) {
		form, _ := c.MultipartForm()  // [] Form
		files := form.File["f1"]  // []*FileHeader
		// 遍历表单中上传的文件
		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("./%d_%s", index,file.Filename)
			// 上传文件到指定的目录
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})

	// 运行服务
	r.Run(":9090")
}
