package main

import "fmt"

//// 求[m，n]之间既不能被7整除也不能被5整除的整数之和，m和n的值由键盘输入。
//例如，如果m和n的值分别为10和20，则计算结果为：106。
func main(){
	var m,n int
	fmt.Println("区间[m，n],m，n为整数,请输入一个m")
	fmt.Scanf("%d",&m)
	fmt.Println("区间[m，n],m，n为整数,请输入一个n")
	fmt.Scanf("%d",&n)
	if m>n{
		fmt.Println("“输入错误")
	}
	var a int
	for i:=m;i<=n;i++ {
		if i%7 != 0 && i%5 != 0{
			a+=i
		}

	}
	fmt.Println("a=",a)//打印输出 在for循环的外面，负责在循环里 会输出多个a的值

}
