package main

import (
	"fmt"
)

// 模拟学生管理系统
func main() {

	c := Class{
		className: "101",
		student:   make([]*Student, 0, 200), // []*Student 指针型数组
	}

	c.addStu("张三", 101, 98)
	c.addStu("李四", 102, 80)
	c.addStu("王五", 103, 100)
	c.addStu("周杰伦", 104, 60)
	c.editStu("admin", 102, 88)
	c.delStu(102)
	c.stuList()

	/*
		// 删除切片中的元素
		seq := []string{"a", "b", "c", "d", "e"}
		index := 2
		// seq[:index]： [a b] ，seq[index+1:]： [d e]
		fmt.Println(seq[:index], seq[index+1:])
		str := append(seq[:index], seq[index+1:]...)
		fmt.Println(str)
	*/

}

// Student 结构体
type Student struct {
	id    int
	name  string
	score int
}

// Class 班级结构体
type Class struct {
	className string
	student   []*Student
}

// 添加学生信息方法
func (c *Class) addStu(name string, id, score int) {
	stu := &Student{
		id:    id,
		name:  name,
		score: score,
	}
	c.student = append(c.student, stu)
	fmt.Printf("add stu %#v\n", stu)
}

// 学生信息列表
func (c *Class) stuList() {
	for _, v := range c.student {
		fmt.Printf("班级：%v,id：%v,姓名：%v，成绩：%v\n", c.className, v.id, v.name, v.score)
	}
}

// 删除学习信息
func (c *Class) delStu(id int) {
	for k, v := range c.student {
		if v.id == id {
			fmt.Printf("del stu: %#v\n", v)
			c.student = append(c.student[:k], c.student[k+1:]...)
		}
	}
}

// 编辑学生信息
func (c *Class) editStu(name string, id, score int) {
	for _, v := range c.student {
		if id == v.id {
			fmt.Printf("before edit stu: %#v\n", v)
			v.id = id
			v.name = name
			v.score = score
			fmt.Printf("after edit stu: %#v\n", v)
		}
	}
}
