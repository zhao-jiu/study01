package main

import "fmt"

// golang函数demo
func main() {
	i, i2 := calc(10, 5)
	fmt.Println(i)
	fmt.Println(i2)

	sum, sub := calc2(5, 5)
	fmt.Println(sum)
	fmt.Println(sub)
}

func calc(x int, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 直接使用定义名字的返回
func calc2(x int, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}
