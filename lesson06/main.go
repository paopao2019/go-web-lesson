package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 嵌套的template的语法知识

func indexFunc(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	msg := "模板嵌套"
	t, err := template.ParseFiles("./index.tmpl","./ul.tmpl")
	if err != nil {
		fmt.Printf("模板解析失败, err: %v\n", err)
	}
	// 渲染模板
	t.Execute(w, msg)


}
func main() {
	http.HandleFunc("/index", indexFunc)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server start failed, err: %v\n", err)
	}

}
