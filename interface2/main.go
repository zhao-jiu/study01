package main

import "fmt"

//一个类型实现多个接口

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

type dog struct {
	name string
}

func (d dog) say() {
	fmt.Printf("%v会旺旺叫。。。\n", d.name)
}

func (d dog) move() {
	fmt.Printf("%v会跑。。。\n", d.name)
}

//多个类型实现同一接口
type car struct {
	brand string
}

// car类型实现Mover接口
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}

func main() {

	var x Sayer
	var y Mover
	var a = dog{name: "旺财"}
	var b = car{brand: "保时捷"}
	x = a
	y = a
	x.say()
	y.move()
	y = b
	y.move()

}
