package main

import "fmt"

//    *          5个空格   1
//   ***             3个空格 3
//  *****               5
// *******           7
//*********       9
// *******
//  *****
//   ***
//    *
	func main() {
		for i := 1; i <= 6; i++ {
			for j := 1; j <= 6-i; j++ {
				fmt.Printf(" ")
			}
			for j := 1; j <= 2*i-1; j++ {
				fmt.Printf("*")
			}
			for j := 1; j <= 6-i; j++ {
				fmt.Printf(" ")
			}
			fmt.Println("")
			//fmt.Println("==")
		}
	}
//     *			5		1	i*2+1
//	  ***			4		3	i*2+1
//   *****			3		5	i*2+1
//  *******			2		7
// *********		1		9
//***********		0		11
//程序有6行 外层循环执行6次

func main(){
	//整体循环控制行
	for i:=0; i<=5;i++  {
		//控制空格个数
		for j:=0;j<5-i;j++{
			fmt.Print(" ")
		}
		//控制星星个数
		for k:=0;k<i*2+1;k++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
