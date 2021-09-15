package main

import "fmt"

//流程控制
func main() {
	// 基本写法
	var score = 60
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 60 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// 特殊写法
	if score := 90; score >= 90 {
		fmt.Println("A")
	} else if score >= 60 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

}
