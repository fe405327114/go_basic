package main

import "fmt"

func main(){
	//aSum是一个函数，不是一个值，不可以用=（1,2）
	aSum(10,20,30,40,50)
}

func aSum(a int,b...int)int{
	sum:=0
	for i,data:= range b{
		fmt.Println("i=",i)
		fmt.Println("data=",data)
        sum+= b[i]

	}
	fmt.Println("sum=",sum)
	return sum
}