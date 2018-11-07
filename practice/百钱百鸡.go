package main

import "fmt"

//一只公鸡值五钱，一只母鸡值三钱，三只小鸡值一钱，现在要用百钱买百鸡，请问公鸡、母鸡、小鸡各多少只？
func main(){
	count:=0
	for a:=0;a<=20;a++{
		for b:=0;b<=33 ;b++  {
			count++
			c:=100-a-b
			if 5*a+3*b+c/3==100 && c%3==0{
				fmt.Printf("公鸡：%d 母鸡：%d 小鸡： %d\n",a,b,c)
			}

		}
	}
	fmt.Println(count)
}