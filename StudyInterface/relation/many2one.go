package relation

import "fmt"

// Mover 接口
//type Mover interface {
//	move()
//}

type doggy struct {
	name string
}

type car struct {
	brand string
}

func (d doggy) move() {
	fmt.Printf("%s会跑\n", d.name)
}

// car类型实现Mover接口
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}

func Many2one()  {
	var x Mover
	a := doggy{name: "旺财"}
	b := car{brand: "林肯"}
	x = a
	x.move()
	x = b
	x.move()
}