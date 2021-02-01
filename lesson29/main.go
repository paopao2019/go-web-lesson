package main

import "fmt"

var baseNum int

// 学习debug  https://zhuanlan.zhihu.com/p/62610785
func sumDemon(x int, y int) (z int) {
	x += baseNum
	y += baseNum
	return x + y
}
func main() {
	a := 10
	b := 10
	fmt.Println(a + b)
	result := sumDemon(10, 20)
	fmt.Printf("和为: %d", result)

}
