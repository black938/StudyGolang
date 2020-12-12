package main

import "fmt"

func structConstructor()  {
	// 构造函数 在前面结构体定义之后的函数 自动生成的NewPerson
	p9:=NewPerson(19,"ww","yy")
	fmt.Println("%#v",p9)
}
func structMethod()  {
	fmt.Println()
}

func structReceiver()  {

}


