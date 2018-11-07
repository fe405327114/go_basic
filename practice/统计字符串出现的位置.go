package main

import "fmt"

//str1 = "hello hello hello world"
//统计"llo"出现的次数:
//统计"wor"出现的位置:
func main(){
	str1 := "hello hello hello world"
	fmt.Println(len(str1))
	s1:=make([]byte,24)
	var i  int
	for i=0;i<len(str1);i++{
	s1[i]=str1[i]
		fmt.Printf("%c",s1[i])
	}
	fmt.Println()
	count:=0
for i,_:=range s1{
	if (int(s1[i])==108)&&(int(s1[i+1])==108)&&(int(s1[i+2])==111){
		count++
	}
}
fmt.Println(count)
}