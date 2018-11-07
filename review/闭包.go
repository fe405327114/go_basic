package main

import "fmt"

//闭包就是 函数作为另一个函数的返回值 进行传递
//通常和匿名函数配合
func demo1()func()int{
	var a int
	return func()int{
		a+=3
		return a
	}
}
func main(){
	b:=demo1()
	for i:=0;i<10;i++{
		fmt.Println(b())
	}

}