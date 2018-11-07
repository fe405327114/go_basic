package main

import (
	"strings"
	"fmt"
)

//3.统计该字符串中每个字母出现的次数
//"ECALIYHWEQAEFSZCVZTWEYXCPIURVCSWTDBCIOYXGTEGDTUMJHUMBJKHFGUKNKN"
//
//
//4.将字符串"aa:zhangsan@163.com!bb:lisi@sina.com!cc:wangwu@126.com"，按照指定的key-value规则存入map中
//key：aa,bb,cc
//value:zhangsan@163.com,lisi@sina.com,wangwu@126.com
func main(){
	var s string="aa:zhangsan@163.com!bb:lisi@sina.com!cc:wangwu@126.com"
    s1:=strings.Split(s,"!")//split函数的返回值为切片类型，有下标i对应v
    fmt.Println(s1)
map1:=make(map[string]string)
	for _,v:=range s1 {
		s2:=strings.Split(v,":")
		map1[s2[0]]=s2[1]
	}
	fmt.Println(map1)
}