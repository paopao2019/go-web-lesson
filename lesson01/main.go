package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//_,_ = fmt.Fprint(w,"Hello world")
	//_,_ = fmt.Fprint(w,"<h1>Hello world</h1><h2>内容</h2>")
	b, _ := ioutil.ReadFile("web.txt")
	fmt.Fprintf(w, string(b))


}
func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http start failed", err)
	}
}
