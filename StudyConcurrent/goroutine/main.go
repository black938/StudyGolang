package main

import (
	"fmt"
	"sync"
)

func main() {
	example()
}

// 简单例子
var wg sync.WaitGroup
func example() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}

func hello(i int) {
	defer wg.Done()
	fmt.Println("Hello Goroutine %s",i)
}
