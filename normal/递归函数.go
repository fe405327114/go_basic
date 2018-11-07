package main

import "fmt"

//求1到n的累加
//func main(){
//	r:=getSum(5)
//	fmt.Println(r)
//
//}
//func getSum(n int)int{
//	if n==1{
//		return 1
//	}
//	return getSum(n-1)+n
//}
//斐波那切数列的算法  前两个数字为1,1，之后数字等与前两个之和
func main(){
	y:=fei(12)
	fmt.Println(y)
}
func fei(i int)int{
	if i==1||i==2{
		return 1
	}
	return fei(i-1)+fei(i-2)
}