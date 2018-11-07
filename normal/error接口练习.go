package main

import "fmt"

func Add(a int,b int)(value int,e error){
	if b==0{
		fmt.Println("除数不能为0")
		return
	}else {
		value=a/b
		return
	}
}
func main(){
	value,e:=Add(10,0)
	if e != nil{
		fmt.Println(e)
	}else {
		fmt.Println(value)
	}
}
