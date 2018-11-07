package main

import (
	"fmt")


//func main(){
//
//
//	var i,sum int
//	for i=0;i<=100 ;i++  {
//
//		sum += i
//	}
//	fmt.Println("0-100的总和为：",sum)


func test(name ...int){


	for i:=0; i<=len(name); i++ {

		fmt.Println(i, name[i])
	}

fmt.Println("======")
}


func test01(haha...int){

	for i,data:=range haha{
		fmt.Println(i,data)


	}
}


//func main(){
//	test01(33,33,33)

//}

func test02()int{
	num1,num2:=89,90
	var sum int
	sum=num1+num2

	return sum

}
func main() {
var a int
a=test02()
test01(33,33,33)

	fmt.Println(a)

}

