package main

import (
	"fmt"
)

// Address 结构体
type Address struct {
	province   string
	city       string
	createTime string
}

type Email struct {
	account    string
	createTime string
}

// User 结构体 嵌套Address
type User struct {
	name string
	age  int
	sex  string
	//address Address
	Address //匿名字段结构体
	Email
}

//----------------go实现结构体继承----------------

//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func main() {
	var user User
	user.name = "小王子"
	user.age = 24
	user.sex = "男"
	user.Address.province = "广东"
	user.city = "广州" // 匿名字段可以省略
	user.account = "10086"
	user.Address.createTime = "2021-09-15 11:27" //指定Address结构体中的CreateTime
	user.Email.createTime = "2021-09-15 11:27"   //指定Email结构体中的CreateTime
	fmt.Printf("user=%#v\n", user)

	fmt.Println("-----------------------------")
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.move()
	d1.wang()

}
