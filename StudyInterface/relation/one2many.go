package relation

import "fmt"

// Sayer 接口
type Sayer interface {
	say()
}
type dog struct {
	name string
}

// 实现Sayer接口
func (d dog) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

// 实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会动\n", d.name)
}

// Mover 接口
type Mover interface {
	move()
}

func One2many() {
	var x Sayer
	var y Mover
	var a = dog{name: "旺财"}
	x = a
	y = a
	x.say()
	y.move()
}
