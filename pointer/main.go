package main

import "fmt"

// 指针
func main() {
	// &取变量的指针
	v := 10
	ptr := &v
	fmt.Println(ptr)

	// *将指针指向的值赋值给变量
	t := *ptr
	fmt.Println(t)

	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)

}
