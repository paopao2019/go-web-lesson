package main


// 原始 net 中的http使用template模板
import (
	"fmt"
	"html/template"
	"net/http"
)

type userS struct {
	Name string
	Age int
	Height float64
}
func sayHello(w http.ResponseWriter, r *http.Request) {

	//1. 定义模板
	//2. 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("解析模板失败 %v", err)
	}
	userM := map[string]interface{} {
		"name": "Andy",
		"age": 18,
		"height": 1.88,

	}
	user := userS{
		Name: "刘德华",
		Age: 17,
		Height: 1.70,
	}

	userList := []string{
		"张三",
		"王五",
		"李四",
	}

	// 3.渲染模板
	t.Execute(w, map[string]interface{} {
		"m": userM,
		"s": user,
		"userList": userList,
	})  // 可以传递任何数据 如map 结构体

	//3. 渲染模板
	//_,_ = fmt.Fprint(w, "<h1>Hello World</h1>")

}
func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http start server failed error is :%v", err)
	}
}
