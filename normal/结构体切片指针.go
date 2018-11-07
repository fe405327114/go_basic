package main

import "fmt"

func main(){
	type stu struct{
		name string
		age int
		score int
	}
	//m1:=make([]stu,2)
	var m1 []stu=[]stu{{"小花",17,78},{"小明",23,56}}
	m1=append(m1,stu{"小刚",23,23})
	fmt.Println(m1)
	p:=&m1
	fmt.Println((*p)[0])

}
