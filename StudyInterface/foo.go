package main

import "fmt"


func (d dog) move() {
	fmt.Println("dog can move")
}

func (c *cat) move()  {
	fmt.Println("cat can move")
}

// 能接受两种类型
func valueReceiver()  {
	var x Mover
	var wangcai = dog{}
	x = wangcai
	x.move()
	var fugui = &dog{}
	x = fugui
	x.move()
}

// 只能接受指针类型
func pointerReceiver(){
	var x Mover
	var mimi = &cat{}
	x = mimi
	x.move()
}
