package main

import "fmt"

//从键盘上输入两个正整数x,y，求它们的最大公约数。

//func main(){
//	var x,y int
//	fmt.Printf("输入两个正整数x，y")
//	fmt.Scanf("%d,%d",&x,&y)
//	if x<=0 || y<=0{
//		fmt.Println("请输入正整数")
//	}
//
//}

//若每步跨2阶，最后剩下1阶；若每步跨3阶，最后剩下2阶；若每步跨5阶，最后剩下4阶；若每步跨6阶，最后剩下5阶；只有每步跨7阶，最后才正好1阶不剩。编程打印这条阶梯共有多少阶。
func main(){
	for x:=0;x<1000;x++{
		if x%2 == 1 && x%3 == 2 && x%5 == 4 && x%6 == 5 && x%7 == 0 {
			fmt.Println(x)
		}
	}
}