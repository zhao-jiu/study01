package main

import "fmt"

// Address 结构体
type Address struct {
	province string
	city     string
}

// User 结构体 嵌套Address
type User struct {
	name string
	age  int
	sex  string
	//address Address
	Address //匿名字段结构体
}

func main() {

	/*u1 := User{
		name: "张三",
		age: 18,
		sex: "男",
		address: Address{
			province: "广东",
			city: "深圳",
		},
	}
	fmt.Println(u1)
	fmt.Printf("u1=%#v\n", u1)
	*/

	var user User
	user.name = "小王子"
	user.age = 24
	user.sex = "男"
	user.Address.province = "广东"
	user.city = "广州" // 匿名字段可以省略
	fmt.Printf("user=%#v\n", user)

}
