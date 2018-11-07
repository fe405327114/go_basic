package main

import "fmt"

func main() {
	var a = [5]int{1234, 24, 546, 75, 1}
	var max, min int
	min = a[0]
	max = a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > max { //如果这里用a[i]和a[i+1]进行比较会越界。需要将i起始位置改为1，a[i]和a[i-1]比较
			max = a[i]
		}
		if a[i] < min {
			min = a[i]
		}
	}
	fmt.Println(max, min)
	//	for i := 0; i < len(a); i++ {
	//		if a[i] < min {
	//			min=a[i]
	//		}
	//	}
	//fmt.Println(min)
}
