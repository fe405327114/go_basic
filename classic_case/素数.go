package main

import "fmt"

//若整数i从左向右读与从右向左读是相同的数，且i为素数，此时称其为回文素数。所谓素数是指只能由1和它本身整除的整数。
//对于偶数位的整数，只有11是回文素数。也就是说，除了11以外，所有的2位整数都不是回文素数。
//所有的4位整数、6位整数、8位整数中也不存在回文素数。
//但是三位回文素数有很多，比如：101、131、151、181、191、313等。本案例要求通过编程求出所有小于1000的回文素数。
func main(){
	//素数
var flag int
	var a int
	for a=11;a<1000;a++ {
		flag=0
		for i:=2;i<a;i++ {
			if a%i==0  {
				flag++
				break
			}
			}
		if flag==0 { //满足此条件则为素数
			if a/100 == 0 {
				if a/10 == a%10 {
					fmt.Println(a) //两位数的素数
				}
			}
			if a/100 == a%10 {
				fmt.Println(a)
			}
		}
	}

}