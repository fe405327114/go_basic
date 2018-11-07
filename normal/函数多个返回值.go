package main

import "fmt"

func main(){
	m,n,z:=test01()
	fmt.Println(m,n,z)
}
func test01()(a int,b int,c int){
	a,b,c=1,2,3
	return
}
