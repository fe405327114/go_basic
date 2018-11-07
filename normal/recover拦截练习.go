package main

import "fmt"

func demo1(i int){
	var a [10]int
	 defer func(){
		err:=recover()
		if  err!=nil{
			fmt.Println(err)
		}
	}()
	a[i]=13
fmt.Println("好啊")
}
func main(){
demo1(11)
}
