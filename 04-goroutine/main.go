package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//test1()
	test2()
}

func test1() {
	//hello()
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second)
}

func test2() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个Goroutine就登记+1
		go hello2(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}

func hello() {
	fmt.Println("Hello Goroutine")
}

func hello2(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}
