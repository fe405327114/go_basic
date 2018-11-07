package main

import "fmt"

//成绩：
//[90,100],优秀
//[80,89],良好
//[70,79],中等
//[60,69],及格
//[0,59],不及格
func main(){
	var score int
	fmt.Scanf("%d",&score)
	if score>=90{
		fmt.Println("优秀")
		if score>95{
			fmt.Println("wa")
		}
	}else if score>=80{
		fmt.Println("良好")
	}else if score>=70{
		fmt.Println("jige")
	}else {
		fmt.Println("bujige")
	}

}

