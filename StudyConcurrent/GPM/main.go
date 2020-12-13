package main

import (
	"fmt"
	"runtime"
	"time"
)

// GPM调度 参考https://www.liwenzhou.com/posts/Go/14_concurrence/
//一个操作系统线程对应用户态多个goroutine
//go程序可以同时使用多个操作系统线程
//goroutine和OS线程是多对多的关系 即m:n
func a() {
	for i := 1; i < 50; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 50; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(8)
	go a()
	go b()
	time.Sleep(time.Second)
}