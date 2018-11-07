package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	randnum:=make([]int,3)
	randnum[0]=rand.Intn(9)+1
	randnum[1]=rand.Intn(10)
	randnum[2]=rand.Intn(10)

	var num int
	usernum:=make([]int,3)
	count := 0
	for{
	for {
		fmt.Println("请输入一个三位数字")
		fmt.Scan(&num)
		if num < 100 || num > 999 {
			fmt.Println("你的输入有误")
		}else {
			break
		}
		}
		usernum[0] = num / 100
		usernum[1] = num / 10 % 10
		usernum[2] = num % 10

	for i:=0;i<3 ;i++  {
		if usernum[i]>randnum[i]{
			fmt.Printf("你输入的第%d位数字太大了\n",i+1)
		}else if usernum[i]<randnum[i] {
			fmt.Printf("你输入的第%d位数字太小了\n",i+1)
		}else {
			fmt.Printf("你输入的第%d位数字是正确的\n",i+1)
			count++
		}
	}
if count==3{
	fmt.Println("你咋这么厉害")
	break
}else {
	count=0
}
	}

}