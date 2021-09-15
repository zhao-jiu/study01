package main

import "fmt"

// 切片demo
func main() {
	var a = []string{"张三", "李四", "王五"} //string数组切片
	fmt.Println(a)
	fmt.Println(len(a))
	var b = []int{1, 5, 7, 8, 9} //int数组切片
	fmt.Println(b)
	var c = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(c)

	sum := sum(b)
	fmt.Println(sum)
	fmt.Println("---------------------------------")
	i := sum2(10, 20, 30, 40, 50)
	j := sum2(10, 20, 30)
	fmt.Println(i)
	fmt.Println(j)
}

// 计算数组中的元素之和
func sum(a []int) int {
	sum := 0
	for _, value := range a {
		sum += value
	}
	return sum
}

// 计算和 传入的值为可变参数
func sum2(x ...int) int {
	//x为切片
	sum := 0
	for _, value := range x {
		sum += value
	}
	return sum
}
