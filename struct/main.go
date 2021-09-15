package main

import "fmt"

// Person 结构体 结构体没有构造函数
type Person struct {
	name string
	age  int
}

// NewPerson 构造函数
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Sleep 睡觉的方法
func (p Person) Sleep() {
	fmt.Printf("%s准备睡觉了！\n", p.name)
}

//Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

// SetAge 方法 指针类型的值可以修改其变量值
func (p *Person) SetAge(newAge int) {
	p.age = newAge
}

func (p *Person) SetName(newName string) {
	p.name = newName
}

func main() {

	//初始化对象
	var p1 Person
	p1.name = "tom"
	p1.age = 18
	fmt.Println("person =", p1)

	var p2 = Person{name: "Jerry", age: 28}
	fmt.Println("person =", p2)

	p3 := Person{name: "Aaron", age: 32}
	fmt.Println("p3 =", p3)

	//匿名结构体
	p4 := struct {
		name string
		age  int
	}{name: "匿名", age: 33}
	fmt.Println("p4 =", p4)

	fmt.Println("-------------------------------")

	// 调用构造函数初始化
	person := NewPerson("小王子", 25)
	fmt.Println(person)
	person.Sleep()
	person.Dream()

	fmt.Println("-------------------------------")
	fmt.Println(person.name, person.age)
	person.SetAge(18)
	person.SetName("迪丽热巴")
	fmt.Println(person.name, person.age)

}
