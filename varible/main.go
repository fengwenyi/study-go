package main

import "fmt"

// 声明变量, 可以放在函数内，也可以放在函数外
// var name string
// var age int
// var ok bool

// 批量声明
// 默认值为该类型的空值
var (
	name string // ""
	age  int    // 0
	ok   bool   // false
)

func main() {
	// 声明变量, 可以放在函数内，也可以放在函数外

	name = "Zhang San"
	age = 18
	// ok = true
	// Go语言变量声明必须使用，不使用编译不过去

	fmt.Printf("name:%s", name)
	fmt.Println() // 打印完自动加一个换行符
	fmt.Println(age)

	// go fmt main.go 自动格式化代码
}
