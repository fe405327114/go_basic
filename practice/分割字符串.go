package main

import (
	"strings"
	"fmt"
)

//将字符串"aa:zhangsan@163.com!bb:lisi@sina.com!cc:wangwu@126.com"，按照指定的key-value规则存入map中
func main(){
	s:="aa:zhangsan@163.com!bb:lisi@sina.com!cc:wangwu@126.com"
	s1:=strings.Split(s,"!")
	fmt.Println(s1)
	m1:=make(map[string]string)
	for i:=0;i<len(s1);i++ {
		s2:=strings.Split(s1[i],":")
		m1[s2[0]]=s2[1]
	}
fmt.Println(m1)
}