package main

import (
	"fmt"
)

//递归函数是自己调用自己  相当于利用函数本身进行for循环，设定条件跳出循环
var a int
var b int=1
func demo2(a int)int{
	if a==0{
		return b
	}
	b=b*a
	a--
	demo2(a)
	return b
}
func main (){
	b:=demo2(5)
	fmt.Println(b)
}
