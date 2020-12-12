package main

import "fmt"

//type Person struct {
//	age        int
//	name, city string
//}
func structConstructor()  {
	// 构造函数 在前面结构体定义之后的函数 自动生成的NewPerson
	p9:=NewPerson(19,"ww","yy")
	fmt.Println("%#v",p9)
}

func structMethod()  {
	p1:=NewPerson(12,"ww","yy")
	p1.Dream()
	p1.SetAge(22)
	fmt.Println(p1)
	var a intMyself
	a.SayHello()
}

// 方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。
// 什么时候应该使用指针类型接收者
// 1 需要修改接收者中的值
// 2 接收者是拷贝代价比较大的大对象
// 3 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
func (p Person) Dream()  {
	fmt.Printf("%s的梦想是学好Golang\n",p.name)
}
// 调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。
// 这种方式就十分接近于其他语言中面向对象中的this或者self
func (p *Person) SetAge(age int)  {
	p.age = age
}
// 接收者的类型可以是任何类型，不仅仅是结构体
type intMyself int
func (i *intMyself) SayHello()  {
	fmt.Println("Hello,im a int")
}


