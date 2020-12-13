package main

import "fmt"

// 空接口是指没有定义任何方法的接口 因此任何类型都实现了空接口
func nilInterface() {
	var x interface{}
	s := "Hello"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
}

// 空接口的应用 接受任意类型的函数参数
func example(a interface{}){
	fmt.Printf("type:%T value:%v\n",a,a)
}

// 空接口的应用2 保存任意值的字典
func example2()  {
	studentInfo := make(map[string]interface{})
	studentInfo["name"] = "woo"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
}

// 类型断言 空接口可以存储任意类型的值 那我们如何获取其存储的具体数据呢
// 一个接口的值由 一个具体类型和具体类型的值 两部分构成 分别叫 动态类型/动态值
// 该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个是一个布尔值
func assertNilInterface(){
	var x interface{}
	x = "Hello 沙河"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}

	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}