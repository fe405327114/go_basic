package main

import "fmt"

//// 求[m，n]之间既不能被7整除也不能被5整除的整数之和，m和n的值由键盘输入。
//例如，如果m和n的值分别为10和20，则计算结果为：106。
func main(){
	var m,n int
	fmt.Println("请输入两个正整数数")

		fmt.Scanf("%d %d",&m,&n)

	if m>n{
		fmt.Println("请输入的m<n")
		return
	}
	sum:=0
	for i:=m;i<n;i++{
		if i%7!=0 && i%5!=0{
			sum+=i
		}
	}
	fmt.Println(sum)
}