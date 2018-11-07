package main

import "fmt"

func main() {
	arr:=[10]int{1,2,3,4,5,6,7,8,9,10}


	for i:=0;i<len(arr)/2;i++{
		arr[i],arr[len(arr)-i-1]=arr[len(arr)-i-1],arr[i]
	}
	////定义起始和结束下标
	//start:=0
	//end:=len(arr)-1
	//
	//for {
	//	if(start>end){
	//		break
	//	}
	//	arr[start],arr[end]=arr[end],arr[start]
	//
	//	start++
	//	end--
	//}

	fmt.Println(arr)
}
