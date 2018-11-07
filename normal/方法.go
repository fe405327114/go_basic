package main

import "fmt"

//函数和方法 的对比
//方法就相当于是结构体中的函数
//1，意义不同
// 函数是具有一段特殊功能的代码，可以多次调用
//方法是某个类的功能，只有指定对象可以调用
//2，语法不同
//函数名不能重复
//方法名字可以一样，但是接收对象不能同时相同
//type 函数名(形参列表)返回值列表{
//	代码体
//}

//type (t Type)方法名(形参列表)返回值列表{
//	代码体
//}
type student struct {
	name string
	age  int
}

func (a student)eat(){
	fmt.Printf("姓名：%s，年龄：%d\n",a.name,a.age)
}
func main(){
	a1:=student{name:"王二狗",age:20}
	a2:=student{name:"李晓华",age:78}
	p:=&a1
	p.eat()
	a1.eat()
	a2.eat()
}

