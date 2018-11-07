package main

import (
	"fmt"
	"strings"
	"sort"
)

//3.统计该字符串中每个字母出现的次数  分别通过map，数组
//"ECALIYHWEQAEFSZCVZTWEYXCPIURVCSWTDBCIOYXGTEGDTUMJHUMBJKHFGUKNKN"

func main(){
	s:="ECALIYHWEQAEFSZCVZTWEYXCPIURVCSWTDBCIOYXGTEGDTUMJHUMBJKHFGUKNKN"
map1:=demo1(s)
	fmt.Println(map1)
demo3(s)

}
//1.利用map
func demo1(s string)map[string]int {
	map1 := make(map[string]int)
	spli := strings.Split(s, "")
	for _,v := range spli {  //v代表的是 字符串中的字符，在map中表示的key
		value,ok := map1[v]//map1的key为每个字符，value为出现的次数
		if ok {
			value++
		} else {
			value = 1
		}
		map1[v]=value
	}
	return map1
}
//2.利用数组   数组的下标为每个字符，对应的值为出现的次数
 func demo2(s string){
 	arr:=[123]int{}
 	fmt.Println(arr)
	 for i:=0;i<len(s);i++{
		 arr[s[i]]++
	 }
	 fmt.Println(arr)
	 for i:=0;i<len(arr);i++{
       if arr[i]!=0{
       	fmt.Printf("%c,%d\n",i,arr[i])
	   }
	 }
 }
 //3.利用数组
 func demo3(s string){
 	arr1:=strings.Split(s,"")
 	sort.Strings(arr1)
 	fmt.Println(arr1)
 	i:=0
 	count:=0
 	for{
 		if arr1[i]==arr1[i+1]{
 			count++
		}else {
			fmt.Println(arr1[i],count)
			count=1
		}
		i++
		if i==len(arr1)-1{
			break
		}
	}
fmt.Println(arr1[len(arr1)-1],count)
 }