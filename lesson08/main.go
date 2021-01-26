package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 更改标识符
// 转义符号 如何 去除默认的脱义

func indexFunc(w http.ResponseWriter, r *http.Request) {
	// 模板定义
	// 模板解析
	t,err := template.New("index.tmpl").
		Delims("{[", "]}").
		ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("模板文件解析失败, err:%v\n", err)
	}

	// 模板渲染
	name := "杰克"
	t.Execute(w, name)

}

// 模拟跨站的问题
func xssFunc(w http.ResponseWriter, r *http.Request) {
	//t, err := template.ParseFiles("./xss.tmpl")

	t,err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string)template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Printf("模板文件解析失败,err :%v", err)
	}

	// 默认是被安全过滤了  有脚本的都会被转义 成正常的字符串显示 html不会执行
	// 那有些需要作为html中的js脚本呢？ 这个时候需要自己定义pipeline函数
	str1 := "<script>alert('呵呵1')</script>"
	// 模拟一段死循环
	str2 := "<script>while(2 >1) {alert(10)} </script>"
	t.Execute(w, map[string]interface{}{
		"str1": str1,
		"str2": str2,
	} )
}

func main() {
	http.HandleFunc("/index", indexFunc)
	http.HandleFunc("/xss", xssFunc)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server start failed, err: %v", err)
	}
}
