package main

import "fmt"

// Nested Struct
type Address struct {
	Province string
	City string
}
type User struct {
	Name string
	Gender string
	Address
}

func nestedAnonymousStruct()  {
	var user User
	user.Name="ww"
	user.Gender="female"
	user.Address.Province ="Anhui"
	user.City= "Hefei" // 注意这里 匿名字段可以省略 只要变量名不冲突
	// 这里是结构体利用匿名字段实现继承语法的关键
	fmt.Println(user)
}

// Inherit
type Animal struct {
	name string
}

func (a Animal) move() {
	fmt.Printf("%s can move \n",a.name)
}

type Dog struct {
	Feet int
	*Animal
}

func (d *Dog) wang()  {
	fmt.Printf("%s can wangwangwang \n",d.name)
}

func inherit(){
	d1:=Dog{
		Feet:   4,
		Animal: &Animal{name: "lucky"},
	}
	d1.move()
	d1.wang()
}