package main

import (
	"errors"
	"fmt"
)

/*
学习函数
*/
func main() {
	sayHello()

	result := intSum(10, 20)
	fmt.Printf("10+20=%d", result) // 30

	fmt.Println()

	result2 := intSum3(1, 2, 3, 4, 5)
	fmt.Printf("1+2+3+4+5=%d", result2) // 15

	fmt.Println()

	result3, result4 := calc(20, 5)
	fmt.Printf("20+5=%d", result3) // 25
	fmt.Println()
	fmt.Printf("20-5=%d", result4) // 15

	fmt.Println()

	result5, result6 := calc2(20, 5)
	fmt.Printf("20+5=%d", result5) // 25
	fmt.Println()
	fmt.Printf("20-5=%d", result6) // 15

	fmt.Println()

	var c calculation               // 声明一个calculation类型的变量c
	c = add                         // 把add赋值给c
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation

	f := add                          // 将函数add赋值给f
	fmt.Printf("type of f:%T\n", f)   // type of f:func(int, int) int
	result7 := f(30, 20)              // 像调用add方法一样调用f
	fmt.Printf("30+20=%d\n", result7) // 50

	result8 := calc3(10, 20, add)
	fmt.Printf("10+20=%d\n", result8) // 30

	result9, err := do("+")
	if result9 != nil {
		fmt.Println(result9(10, 50))
	}
	fmt.Println(err)

	// 匿名函数
	anonymousFunc()

	// 闭包
	fmt.Println("闭包测试-start")
	var closureFunc = adder()
	fmt.Println(closureFunc(10)) // 10
	fmt.Println(closureFunc(30)) // 40
	fmt.Println(closureFunc(50)) // 90
	fmt.Println(closureFunc(70)) // 160
}

func sayHello() {
	fmt.Println("Hello Go") // Hello Go
}

// 求和函数
func intSum(x int, y int) int {
	return x + y
}

// 也可以简写
func intSum2(x, y int) int {
	return x + y
}

// 可变参数
// 注意：可变参数一定要放在所有参数的最后面
func intSum3(x ...int) int {
	fmt.Println(x) // x是一个切片 // [1 2 3 4 5]
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// 多参数返回，必须用英文括号()包裹起来
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 也可以定义返回变量名
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 空
func someFunc(x string) []int {
	if x == "" {
		return nil // 没必要返回 []int{}
	}
	return []int{1, 2, 3}
}

// 变量作用域

// 全局变量

// 局部变量

// 函数类型与变量
// 定义函数类型
type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 函数作为参数
func calc3(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

// 匿名函数
func anonymousFunc() {

	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	// 通过变量调用匿名函数
	add(10, 20)

	// 自执行函数，匿名函数定义完加括号()，直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 50)
}

// 闭包
// 闭包 = 函数 + 引用环境
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
