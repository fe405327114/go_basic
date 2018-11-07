package main

import "fmt"

//按下面的公式求sum的值。
//sum = m + (m+1) + (m+2) + (m+3) + …… + (n-1) + n 例如，如果m和n的值分别为1和100，则计算结果为5050。
func main(){
	var m,n int
	var sum int
	fmt.Println("输入m和n的值")
	fmt.Scanf("%d,%d",&m,&n)
	for i:=m;i<=n;i++{
		sum+=i
	}
	fmt.Printf("m=%d\nn=%d\n",m,n)
	fmt.Println("sum=",sum)
}