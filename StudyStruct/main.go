package main

import (
	"fmt"
	"unsafe"
)

type NewInt int  // type of it: main.MyInt 重新整了一个
type MyInt = int // type of it: int 起别名 并没有起新类型
//type byte uint8
//type rune int32

type Person struct {
	age        int
	name, city string
}

func NewPerson(age int, name string, city string) *Person {
	return &Person{age: age, name: name, city: city}
}
func main() {
	base()
	memoryAlign()
	problem()
	// foo.go
	structConstructor()
	structMethod()
	// bar.go
	nestedAnonymousStruct()
	inherit()
	setSliceMap2Struct()
	//json.go
	tryJSON()

}

func base() {
	//基本实例化
	var p1 Person
	p1.name = "WangSir"
	p1.city = "Nanjing"
	p1.age = 22
	//匿名结构体
	var user struct {
		name string
		age  int
	}
	user.name = "WangMadam"
	user.age = 20
	//创建结构体指针
	var p2 = new(Person)
	p2.age = 20
	p2.name = "temp"
	fmt.Println("%T", p2)
	fmt.Println("p2 = %#v", p2)
	//取结构体的地址实例化
	p3 := &Person{}
	p3.age = 30
	p3.name = "qimi"
	//p3.name = "七米"其实在底层是(*p3).name = "七米"，这是Go语言帮我们实现的语法糖。

	//键值对初始化
	p5 := Person{
		age:  20,
		name: "wangdachui",
		city: "maanshan",
	}
	fmt.Println("%#v", p5)
	//值的列表初始化
	p6 := Person{21, "wangxiaochui", "nanjing"}
	fmt.Println("%#v", p6)

}

func memoryAlign() {
	//进阶 体会内存对齐带来的影响 https://segmentfault.com/a/1190000017527311
	//根据当前类型大小来确定对齐值 内存存放的偏移量必须是对齐值的整数倍
	type Part1 struct {
		a bool
		b int32
		c int8
		d int64
		e byte
	}
	// axxx|bbbb|cxxx|xxxx|dddd|dddd|exxx|xxxx
	type Part2 struct {
		e byte
		c int8
		a bool
		b int32
		d int64
	}
	// ecax|bbbb|dddd|dddd
	part1 := Part1{}
	part2 := Part2{}
	fmt.Printf("part1 size: %d, align: %d\n", unsafe.Sizeof(part1), unsafe.Alignof(part1))
	fmt.Printf("part2 size: %d, align: %d\n", unsafe.Sizeof(part2), unsafe.Alignof(part2))
}

func problem() {
	// 一道面试题 求输出 https://studygolang.com/articles/9701
	type student struct {
		name string
		age  int
	}
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

}
