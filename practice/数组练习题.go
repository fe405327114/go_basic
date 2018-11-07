package main

import (
	"fmt"
	"math/rand"
	"time"
)

//数据是四位的整数，每位数字都加上5,然后用和，除以10的余数代替该数字，
// 再将第一位和第四位交换，第二位和第三位交换。



//	var arrt[4]int=[4]int{0,1,2,3}  //注意数组的全部数据初始化的定义格式
//	var sum int
//
//	for i:=0;i<len(arrt);i++{
//		sum=sum+arrt[i]+5
//	}
//	for i:=0;i<len(arrt) ;i++ {
//		arrt[i]=sum%10
//	}
//	arrt[0],arrt[1]=arrt[3],arrt[2]
//
//	fmt.Println(arrt)
//
//}


//6.给定一个整型数组，长度为10。数字取自随机数。
//7.将作业6的数组进行从大到小排序。

	//冒泡排序
	func main() {
	rand.Seed(time.Now().UnixNano())
	var ae [10]int
	for i := 0; i < 10; i++ {
		//把取得的随机数重复数字去掉
			for {
				count:=0
				x := rand.Intn(100)
				for j := 0; j < i; j++ {
					if x == ae[j] {
						count++
						break
					}
					}
					ae[i] = x
				if count ==0{
					break
				}
			}
	}
		fmt.Println(ae)


	for m := 0; m < len(ae)-1; m++ {
		for n := 0; n < len(ae)-m-1; n++ {
			if ae[n] > ae[n+1] {
				ae[n], ae[n+1] = ae[n+1], ae[n]
			}
		}
	}
	fmt.Println("排序后：", ae)
}