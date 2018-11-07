package main

import (
	"math/rand"
	"time"
	"fmt"
)

//6.给定一个整型数组，长度为10。数字取自随机数，并且不能重复
//7.将作业6的数组进行从大到小排序。
func main(){
	var app [10]int
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<10 ;i++  {
		var x int
		for { //这里for的意思是，下面的break 跳出下面循环后，并未轮到上层for的i++操作，要继续
			//取随机数进行本次i的循环，但是这里for需要一个条件去跳出取随机数，那就是app[j]从未==过X
			//使用计数器
			count := 0
			x = rand.Intn(10)
			for j := 0; j < i; j++ {
				if app[j] == x {
					count++
					break
				}
			}
			if count == 0 {
				break
			}
		}
		app[i]=x
	}

	fmt.Println(" 排序前",app)
//冒泡排序
	for i:=0;i<len(app)-1 ;i++  {
		for j:=0;j<len(app)-i-1 ; j++ {
			if app[j]>app[j+1]{
				app[j],app[j+1]=app[j+1],app[j]
			}
		}
	}

	fmt.Println("排序后",app)


}