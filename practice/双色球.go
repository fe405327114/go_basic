package main

import (
	"math/rand"
	"time"
	"fmt"
)

//福利彩票  双色   红球有6个  1-33  不能重复  蓝球 1-16 随机一个
//rand.Intn(33)+1//1-33
//func main(){
//	var arr [6]int
//	rand.Seed(time.Now().UnixNano())
// for i:=0;i<6;i++{
//	 var x int
//	 for {
//		 count:=0
//		 x = rand.Intn(33)+1
//		 for j := 0; j < i; j++ {
//			 if arr[j] == x {
//				 count++
//				 break
//			 }
//		 }
//		 if count==0{
//		 	break
//		 }
//	 }
//	 arr[i]=x
// }
// fmt.Println(arr)
//}
func main() {
	var arr [6]int
	rand.Seed(time.Now().UnixNano())
	temp := rand.Intn(33) + 1
for i:=0;i<6;i++{
	for j:=0;j<i ;j++ {
		if arr[j] == temp{
			temp = rand.Intn(33) + 1
			j=-1
			continue
		}
		}
		arr[i]=temp
	}
	fmt.Println(arr,"+",rand.Intn(16)+1)
}
