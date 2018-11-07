package main

import "fmt"

func main(){

	a:= []int{0,1,2,3,4,5,6,7,8,9}
	s1:=a[:6]
	fmt.Printf("len=%d,cap=%d\n",len(s1),cap(s1))
	s1=a[6:]
	fmt.Printf("%d\n%d\n",len(s1),cap(s1))
	s1=a[2:5]
	fmt.Printf("chang:%d\nrong:%d",len(s1),cap(s1))

}
