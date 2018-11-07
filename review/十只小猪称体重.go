package main

import "fmt"

//找到最终和最轻的
func main() {
	arr := [10]int{23, 435, 46, 65, 7, 68, 8, 86, 6, 100}
	max ,min:= 0,0
	for i := 0; i < len(arr)-1; i++ {
		if max < arr[i] {
			max = arr[i]
		}else{
			min=arr[i]
		}
	}
	fmt.Println(max,min)
}