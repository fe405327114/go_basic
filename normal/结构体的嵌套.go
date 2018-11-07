package main

import (
	"fmt"
)

type book struct {
	书名 string
	作者年龄 int
}

type student struct {
	name string
	age int
	book book
}
func main(){
	a1:=book{书名:"三国演义",作者年龄:109}
	//a1.书名="三国演义"
	//a1.作者年龄=109
	s1:=student{age:34,name:"王二狗",book:a1}
	//s1.age=34
	//s1.name="王二狗"
	//s1.book=a1
	fmt.Println("姓名：",s1.name,"年龄：",s1.age,"看的书有:" ,s1.book.书名,"作者年龄是：",a1.作者年龄)
	fmt.Println(a1)

	type student1 struct {
		name string
		age int
		book []book

	}
s2:=student1{}
	a2:=book{书名:"红楼梦",作者年龄:135}
	s2.name="白雪"
	s2.age=24
	s2.book=make([]book,0,10)
	s2.book=append(s2.book,a2,a1)

fmt.Println(s2)
}
