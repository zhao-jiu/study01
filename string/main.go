package main

import (
	"fmt"
	"strings"
)

//字符串操作
func main() {
	//字符串常用操作
	s1 := "hello"
	s2 := "你好北京"
	fmt.Println(len(s1))

	//拼接字符串
	fmt.Println(s1 + s2)
	s3 := fmt.Sprintf("%s - %s", s1, s2)
	fmt.Println(s3)

	//分割字符串
	s4 := "who are you are"
	split := strings.Split(s4, " ") //切割成数组
	fmt.Printf("%s - %s - %s\n", split[0], split[1], split[2])

	//查看类型
	fmt.Printf("%T\n", split)

	//是否包含
	contains := strings.Contains(s4, "you")
	fmt.Println(contains)

	//字符串出现的位置
	fmt.Println(strings.Index(s4, "are"))
	fmt.Println(strings.LastIndexAny(s4, "are"))

	//join操作
	s5 := []string{"how", "do", "you", "do"}
	fmt.Println(s5)
	fmt.Println(strings.Join(s5, "-")) //使用-连接字符串切片

}
