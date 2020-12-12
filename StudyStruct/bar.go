package main

import "fmt"

// Nested Struct
type Address struct {
	Province string
	City     string
}
type User struct {
	Name   string
	Gender string
	Address
}

func nestedAnonymousStruct() {
	var user User
	user.Name = "ww"
	user.Gender = "female"
	user.Address.Province = "Anhui"
	user.City = "Hefei" // 注意这里 匿名字段可以省略 只要变量名不冲突
	// 这里是结构体利用匿名字段实现继承语法的关键
	fmt.Println(user)
}

// Inherit
type Animal struct {
	name string
}

func (a Animal) move() {
	fmt.Printf("%s can move \n", a.name)
}

type Dog struct {
	Feet int
	*Animal
}

func (d *Dog) wang() {
	fmt.Printf("%s can wangwangwang \n", d.name)
}

func inherit() {
	d1 := Dog{
		Feet:   4,
		Animal: &Animal{name: "lucky"},
	}
	d1.move()
	d1.wang()
}

// 因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意
type Per struct {
	name   string
	age    int8
	dreams []string
}

// 错误写法 因为传过去的参数是一个指针 所以这边set去了这个参数的指针 数据在传来的参数那里 很危险
//func (p *Per) SetDreams(dreams []string) {
//	p.dreams = dreams
//}

// 正确写法 要开辟一块新空间 然后copy过来
func (p *Per) SetDreams(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

func setSliceMap2Struct() {
	p1 := Per{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)
	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams)
}
