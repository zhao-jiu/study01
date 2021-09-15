package main

import (
	"fmt"
	"math"
)

//数据类型
func main() {
	//十进制
	a := 10
	fmt.Printf("%d\n", a) //十进制
	fmt.Printf("%b\n", a) //二进制

	//八进制
	m := 077
	fmt.Printf("%d\n", m)
	fmt.Printf("%o\n", m) //打印八进制

	//十六进制
	f := 0xff
	fmt.Println(f)
	fmt.Printf("%x\n", f) //打印十六进制

	//浮点型
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.3f\n", math.Pi) //3f保留三位小数

	//布尔类型
	var result bool
	result = true
	fmt.Println(result)

	//字符串
	//多行字符串
	str := ` 
		a
		b
		c
		d`
	fmt.Println(str)

	str1 := "hello beijing"
	str2 := "你好北京！"
	fmt.Println(str1)
	fmt.Println(str2)

}
