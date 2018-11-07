package main

import (
	"fmt"
)

func main(){
	m1:=map [int]string{1:"a",2:"b",3:"c"}
	m2:=make(map[int]string)
	fmt.Println("m1[1]=",m1[1])
	fmt.Println("m2[1]=",m2[1])
	s1:=[]int{1,2,3,4}
	s2:=make([]int,4,10)
	fmt.Printf("len(s1)=%d\ncap(s1)=%d\n",len(s1),cap(s1))
	fmt.Printf("len(s2)=%d\ncap(s2)=%d\n",len(s2),cap(s2))
}
