package main

import (
	"fmt"
	"strings"
)

// map操作demo
func main() {
	// 初始化map
	scoreMap := make(map[string]int, 8)
	// 赋值操作
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Println(scoreMap["张三"])
	fmt.Printf("type of scoreMap:%T\n", scoreMap)

	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	// v -> 值  ok -> bool
	v, ok := scoreMap["小明"]
	fmt.Println(ok)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("没有找到对应的结果")
	}

	// 再初始化map的时赋值
	userInfo := map[string]string{
		"username": "张三",
		"password": "123456",
	}
	fmt.Printf("userInfo:%v\n", userInfo)

	// 遍历map key value
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	//只遍历key值
	for k := range scoreMap {
		fmt.Printf("key:%v\n", k)
	}

	// 统计字符串中相同值出现的个数
	str := "how do you do"
	fmt.Println("----------------------")
	countMap := count(str)
	fmt.Printf("countMap：%v\n", countMap)

	fmt.Println("----------------------")
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("s:%v\n", s)
	m["q1mi"] = s
	fmt.Printf("m:%v\n", m)
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])

}

//	统计字符串中相同值出现的个数
//	return map[string]int
func count(str string) map[string]int {
	split := strings.Split(str, " ")
	countMap := make(map[string]int, 8)
	// 遍历切片
	for _, key := range split {
		_, isReal := countMap[key]
		// 值不存在 初始化数量
		if !isReal {
			countMap[key] = 1
		} else {
			// 当值存在，数量加1
			countMap[key] += 1
		}
	}
	return countMap
}
