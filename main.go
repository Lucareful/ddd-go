package main

import (
	"fmt"
	"time"
)

func main() {
	// main函数中goroutine如果结束，子goroutine不会再执行
	// 这里你把 main（特殊的goroutine） 当作一个 "进程" 来理解，子 goroutine 就是 main 中的 "线程"
	// "进程"退出了那么"线程"不管有没有执行完都会退出
	fmt.Println("main函数开始")
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("main中子goroutine结束")
	}()
	fmt.Println("main函数结束")
}

/*
out:
	main函数开始
	main函数结束
*/
