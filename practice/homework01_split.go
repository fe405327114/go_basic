package main

import (
	"strings"
	"fmt"
)

func main()  {
	/*
	2.将字符串"aa:zhangsan@163.com!bb:lisi@sina.com!cc:wangwu@126.com"，按照指定的key-value规则存入map中
	key：aa,bb,cc
	value:zhangsan@163.com,lisi@sina.com,wangwu@126.com
	 */
	 str:="aa:zhangsan@163.com!bb:lisi@sina.com!cc:wangwu@126.com"
	 s1:=strings.Split(str,"!")
	 //fmt.Println(s1)
	map1 := make(map[string]string)
	 for _,v:=range s1{
	 	//fmt.Println(i, v)
	 	s2:=strings.Split(v,":")
	 	map1[s2[0]] = s2[1]
	 }
	 fmt.Println(map1)


}
