package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
/*
"ECALIYHWEQAEFSZCVZTWEYXCPIURVCSWTDBCIOYXGTEGDTUMJHUMBJKHFGUKNKN"
"E"-->2
"C"-->1
"A"-->2

A：65，Z：90
a：97，z：122
//
 */
 str := "ECALIYHWEQAEFSZCVZTWEYXCPIURVCSWTDBCIOYXGTEGDTUMJHUMBJKHFGUKNKN"
 map1:=count1(str)
 fmt.Println(map1)

 count2(str)

 count3(str)

}

func count1(str string) map[string]int{
	map1 :=make(map[string]int)
	arr := strings.Split(str,"")
	for _,v:=range arr{
		value,ok:=map1[v]
		if ok{
			value++
		}else{
			value = 1
		}
		map1[v] =value
	}
	return map1
}

func count2(str string){
	arr :=[123]int{}
	fmt.Println(arr)
	for i:=0;i<len(str);i++{
		arr[str[i]]++
	}
	fmt.Println(arr)
	for i:=0;i<len(arr);i++{
		if arr[i] != 0{
			fmt.Printf("%c,%d\n",i, arr[i])
		}
	}
}

func count3(str string){
	arr := strings.Split(str,"")
	fmt.Println(arr)
	fmt.Printf("%T\n",arr)
	sort.Strings(arr)//排序
	fmt.Println(arr)
	index := 0
	count :=1
	for {
		if arr[index] == arr[index+1]{
			count++
		}else{
			fmt.Println(arr[index], count)
			count = 1
		}
		index++
		if index == len(arr)-1{
			break
		}
	}
	fmt.Println(arr[len(arr)-1], count)

}
