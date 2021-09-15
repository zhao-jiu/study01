package main

import (
	"fmt"
	"github.com/q1mi/hello"
)

func main() {
	var name string
	name = "小王子"
	fmt.Println(name)
	var a = 10
	fmt.Println(a)
	fmt.Println("Hello let's Go!")
	hello.SayHi()

	var str = "aaa"
	fmt.Println(str)

	//批量声明变量
	var (
		age   int
		sex   string
		email string
	)
	age = 10
	sex = "男"
	email = "123456@qq.com"
	fmt.Println(age)
	fmt.Println(sex)
	fmt.Println(email)

	//常量
	const pi = 3.1415
	const e = 2.7182

	f := pi * e
	println(f)

	str1 := `aa,bb,cc`
	println(str1)

	str2 := "hello沙河小王子"
	fmt.Println(len(str2))

}
