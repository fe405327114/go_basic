package main

import "fmt"

func main(){
	type student1 struct {
		id int
		name string
		score int
	}
	m1:=make(map[int]student1)
	m1[101]=student1{1,"狗狗",23}
	m1[203]=student1{2,"大柴",235}
	for k,v:=range m1{
		fmt.Println(k,v)
	}
	m2:=make(map[int][]student1)
	m2[102]=[]student1{student1{3,"大黄",243}} // error 切片未初始化之前是空的
	//或者用append进行切片元素的添加
	m2[305]=append(m2[305],student1{4,"小狗",34},student1{5,"大鱼",35})
	m2[405]=append(m2[405],student1{6,"鱼鱼",24},student1{7,"柴柴",35})
	for k,v:=range m2{
		fmt.Println(k,v,"===========")
		m2[305][0].name="白雪"
		for i, data:=range v{
			fmt.Println(k,i,data)
		}
	}
}