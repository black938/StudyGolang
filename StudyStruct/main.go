package main

import "fmt"

type Student struct {
	id   int
	name string
}
type NewInt int  // type of it: main.MyInt 重新整了一个
type MyInt = int // type of it: int 起别名 并没有起新类型
type byte uint8
type rune int32

func main() {
	xiaoming := Student{id: 1, name: "xiaoming"}
	fmt.Printf("%#v", xiaoming)
	type foo struct {
		id int
	}

}
