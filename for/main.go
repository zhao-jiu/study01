package main

import "fmt"

//for循环
func main() {

	// 九九乘法表
	for i := 0; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}

}

func add(a int, b int) int {
	return a + b
}

func reduce(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}
