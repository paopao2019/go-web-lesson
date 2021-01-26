package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexFunc(w http.ResponseWriter, r *http.Request) {
	msg := "Index模板文件"
	// 解析模板文件
	t,err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("解析模板文件失败, err %v\n", err)
		return
	}
	// 渲染模板
	t.Execute(w, msg)
}

func homeFunc(w http.ResponseWriter, r *http.Request) {
	msg := "Home模板文件"
	// 解析模板文件
	t,err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("解析模板文件失败, err %v\n", err)
		return
	}
	// 渲染模板
	t.Execute(w, msg)
}


//模板继承的使用
func index2Func(w http.ResponseWriter, r *http.Request) {
	msg := "Index2模板文件"
	// 解析模板文件  经测试 解析文件的前后顺序没有要求。 但是最好是先有base才有被继承的文件
	t,err := template.ParseFiles("./templates/base.tmpl","./templates/index2.tmpl")
	//t,err := template.ParseFiles("./templates/index2.tmpl","./templates/base.tmpl")
	if err != nil {
		fmt.Printf("解析模板文件失败, err %v\n", err)
		return
	}
	// 渲染模板  因为解析了多个模板 所以要用ExecuteTemplate的方式 指定模板  注意名字是basefilename 不要写成 "./templates/index2.tmpl"
	t.ExecuteTemplate(w, "index2.tmpl", msg)

}

func home2Func(w http.ResponseWriter, r *http.Request) {
	msg := "Home模板文件"
	// 解析模板文件
	t,err := template.ParseFiles("./templates/base.tmpl","./templates/home2.tmpl")
	if err != nil {
		fmt.Printf("解析模板文件失败, err %v\n", err)
		return
	}
	// 渲染模板
	//t.Execute(w, msg)
	t.ExecuteTemplate(w, "home2.tmpl", msg)
}

func main() {
	http.HandleFunc("/index", indexFunc)
	http.HandleFunc("/home", homeFunc)
	http.HandleFunc("/index2", index2Func)
	http.HandleFunc("/home2", home2Func)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server start err, err:%v\n", err)
	}
}
