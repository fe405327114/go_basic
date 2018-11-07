package main

import "fmt"

//递归函数是自己调用自己
func demo1 (n int)int{
	if n==1||n==2{
		return 1
	}
	return demo1(n-1)+demo1(n-2)
}
func main (){
	r:=demo1(12)
	fmt.Println(r)
}
