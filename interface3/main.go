package main

import (
	"fmt"
)

// 并一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。

type WashingMachine interface {
	wash()
	dry()
}

type dryer struct{}

func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

//  接口嵌套

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}

type cat struct{}

func (c cat) say() {
	fmt.Println("喵喵喵")
}
func (c cat) move() {
	fmt.Println("猫会动")
}

func main() {
	h := haier{dryer{}}
	h.dry()
	h.wash()

	fmt.Println("--------------------")

	var a animal
	c := cat{}
	a = c
	a.say()
	a.move()

	fmt.Println("--------------------")

	// 空接口作为map值 可以存储任意类型
	var studentInfo = make(map[string]interface{}, 8)
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)

}
