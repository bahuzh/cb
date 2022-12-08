package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func (stu *Student) GetUserInfo() {
	fmt.Printf("姓名：%s \n年龄: %d", stu.Name, stu.Age)
}

func main() {
	student := Student{"zz", 20}
	student.GetUserInfo()
}
