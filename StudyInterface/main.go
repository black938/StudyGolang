package main

import (
	"fmt"
	r "Study/StudyInterface/relation"
)

// Go的interface是一种类型 一种抽象的类型 interface是一组method的集合
// Go提倡面向接口编程 当你看到一个接口类型的值时，你不知道它是什么，唯一知道的是通过它的方法能做什么
// 一个对象只要实现了接口的全部方法 那就实现了这个接口 接口就是一个需要实现的方法列表

type Mover interface {
	move()
}
type Sayer interface {
	say()
}

type dog struct {
}

type cat struct {
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

func main() {
	// 接口类型变量能够存储所有实现了该接口的实例
	var x Sayer
	a := dog{}
	b := cat{}
	x = a
	x.say()
	x = b
	x.say()


	r.Many2one()
	r.One2many()
	r.NestedInterface()
	nilInterface()
}

// 来体会一下这个 _ 的妙用 ：  确保RouterGroup实现了接口IRouter
// 摘自gin框架routergroup.go
//type IRouter interface{ ... }
//type RouterGroup struct { ... }
//var _ IRouter = &RouterGroup{}

// 类型和接口的关系
// 一个类型实现多个接口 多个类型实现同一接口