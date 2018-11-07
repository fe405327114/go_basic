package main

import "fmt"

//按下面的公式求sum的值。
//sum = 1 - 2 + 3 - 4 + 5 - 6 + …… + 99 - 100
func main(){

	var sum =0
	for i:=1;i<=100;i++{
		if i%2==0{
				sum-=i
		}else {
        sum+=i
        }
	}
	fmt.Println("sum=",sum)
}