package main

import (
	"fmt"
)

//数组demo
func main() {
	//var a[3]int
	//var b[4]int
	//fmt.Println(a)
	//fmt.Println(b)

	//数组初始化
	var city1 = [4]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(city1)

	//不定长度的数组
	var city2 = [...]string{"北京", "上海", "广州", "深圳", "杭州", "成都"}
	fmt.Println(city2)

	//使用索引值初始化数组
	var langArray = [...]string{1: "Golang", 3: "Java"}
	fmt.Println(langArray)

	//数组的遍历
	//1.for循环遍历
	//for i:=0;i<len(city2);i++ {
	//	fmt.Println(city2[i])
	//}

	//2.for rang遍历
	//for index, value := range city2 {
	//	fmt.Println(index,value)
	//}

	//只获取数组的元素值
	//for _, value := range city2 {
	//	fmt.Println(value)
	//}

	//二维数组
	cityArray := [...][2]string{
		{"北京", "天津"},
		{"上海", "杭州"},
		{"重庆", "成都"},
		{"广州", "深圳"},
	}
	//遍历二维数组
	for _, v1 := range cityArray {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	//数组在go语言中是值类型  通过下面的例子发现赋值没有改变数组的元素值
	a := [3]int{1, 2, 3}
	fmt.Println(a)
	f1(a)
	fmt.Println(a)

	b := []int{1, 3, 5, 7, 9}
	i := sum(b)
	fmt.Println(i)

}

func f1(a [3]int) {
	a[0] = 100
}

func sum(a []int) int {
	sum := 0
	for _, value := range a {
		sum += value
	}
	return sum
}
