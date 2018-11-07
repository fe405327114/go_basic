package main

import "fmt"

//首先在屏幕上输入一句话，每个单词之间用一个空格隔开，要求第一个字符和最后一个字符都不能为空格；
//然后统计出这句话的单词数量，并把结果输出到屏幕上。
func main(){
	//s:=make([]string,50)
	count:=0
	//fmt.Println("请输入一句话：")
	//fmt.Scan(&s)
	//fmt.Println(s)
	s:="s s d"
	 s1:=make([]string,100)
	for i:=0;i<len(s);i++ {
		s1[i]=string(s[i])
	}
	for i,_:=range s1{
		if s1[i]==""{
			count++
		}
	}
	fmt.Println(count+1)
}