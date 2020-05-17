package main

import "fmt"

/*
学习常量
相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。
常量的声明和变量声明非常类似，只是把var换成了const，常量在定义的时候必须赋值。
*/
func main() {

	/*
		声明了pi和e这两个常量之后，在整个程序运行期间它们的值都不能再发生变化了。
	*/
	const pi = 3.1415
	const e = 2.7182

	// 多个常量也可以一起声明
	const (
		n1 = 1
		n2 = 2
		n3 = 3
	)

	// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。
	const (
		x = 1
		y
		z
	)
	fmt.Println("x=", x)
	fmt.Println("y=", y)
	fmt.Println("z=", z)

	fmt.Println("-- iota ---")

	// iota
	const (
		x1 = iota
		x2
		x3
		x4
	)

	fmt.Println("x1=", x1)
	fmt.Println("x2=", x2)
	fmt.Println("x3=", x3)
	fmt.Println("x4=", x4)
}
