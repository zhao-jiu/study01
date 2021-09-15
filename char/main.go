package main

import "fmt"

//字符
func main() {

	// byte int8的别名
	// rune int32的别名
	var a1 byte = 'c'
	var a2 = 'c'
	fmt.Println(a1, a2)
	fmt.Printf("a1:%T  a2:%T\n", a1, a2)

	s1 := "hello world"
	for _, r := range s1 {
		fmt.Printf("%c\n", r)
	}

}
