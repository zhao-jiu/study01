package main

import "fmt"

func main() {
	// 判断单个值
	score := 1
	switch score {
	case 1:
		fmt.Println("A")
	case 2:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	default:
		fmt.Println("无效的输入")
	}

	//一次判断多个值
	num := 2
	switch num {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	}

	// case带表达式
	age := 18
	switch {
	case age >= 18:
		fmt.Println("可以上网了！")
	case age < 18:
		fmt.Println("回家洗洗睡吧")
	}

}
