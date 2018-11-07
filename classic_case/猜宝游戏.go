package main

import (
	"math/rand"
	"time"
	"fmt"
)

//左手有一个硬币，经过有限次的随机交换后，打印出硬币在哪只手里
//思路：设计两个函数，一个交换形参，一个交换实参，交换次数随机产生
func case1(l,r int){
	l,r=r,l
}
func case2(l,r *int){
	*l,*r=*r,*l
}
func main(){
	var a  int
	l:=1
	r:=0
	rand.Seed(time.Now().UnixNano())
	a=rand.Intn(100)%2
	if a==1 {
		case1(l,r)
		fmt.Println(l,r)
	}else {
		case2(&l,&r)
		fmt.Println(l,r)
	}

}
