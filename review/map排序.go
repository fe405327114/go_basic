package main

import (
	"sort"
	"fmt"
)

func main(){
	//m2:=make(map[string]string)//用这种方法必须分别对每一个K值进行赋值操作
	m2:=map[string]string{"d":"小明","a":"小红","b":"小花"}
	s2:=make([]string,0,len(m2))
for k,_:=range m2{
	s2=append(s2,k)
}
sort.Strings(s2)
fmt.Println(s2)
for _,v:=range s2{
	fmt.Println(v,m2[v])
}
}