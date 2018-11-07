package main

import "fmt"

func main(){
f:=func (a int ,b int )int{
	return a+b
}
y:=func (){
	fmt.Println("hello world")
}
	v:=f(10,20)
	fmt.Println(v)

	y()
}

