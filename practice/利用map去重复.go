package main

import (
	"strings"
	"fmt"
)

func main(){
	var str string = "sdsdfgfghjhj"
	s1:=strings.Split(str,"")
	m1:=make(map[int]string)
	var j int
	for _,v:=range s1{
		count:=0
		for j=0;j<len(m1);j++ {
			if v==m1[j]{
				count++
				break
			}
		}
		if count==0{
			m1[j]=v
		}
	}
	fmt.Println(m1)

}
