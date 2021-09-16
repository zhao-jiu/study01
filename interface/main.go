package main

import "fmt"

// 说话的接口
type sayer interface {
	say()
}

// 实现
func (d dog) say() {
	fmt.Println("汪汪汪。。。。")
}

// 实现
func (c cat) say() {
	fmt.Println("喵喵喵。。。。")
}

type dog struct{}
type cat struct{}

// ---------------------------------

//值接收者和指针接收者实现接口的区别

type Mover interface {
	move()
}

func (d dog) move() {
	fmt.Println("狗会跑。。。")
}

//func (d *dog) move() {
//	fmt.Println("狗会跑。。。")
//}

// 接口demo
func main() {

	a := dog{}
	b := cat{}
	//dog.say()
	//cat.say()

	var x sayer // 定义接口变量
	x = a       //可以把dog实例直接赋值给x
	x.say()
	x = b //可以把cat实例直接赋值给x
	x.say()

	fmt.Println("-------------")

	var y Mover

	var wangcai = dog{} // 旺财是dog类型
	y = wangcai         // x可以接收dog类型  当传入的类型为指针类型时 这个赋值不可以
	var fugui = &dog{}  // 富贵是*dog类型
	y = fugui           // x可以接收*dog类型
	y.move()

}
