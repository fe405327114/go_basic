package main

import (
	"fmt"
)

//dsfjgjfgdsnnjgnfdjngjdfngjdnnrege
func main(){
	s:="dsfjgjfgdsnnjgnfdjngjdfngjdnnrege"
	//s1:=strings.Split(s,"")
     sli:=make([]int,123)
    for i:=0;i<len(s);i++{
	sli[s[i]]++
	}
	fmt.Println(sli)
	for i,_:=range sli{
		if sli[i]!=0 {
			fmt.Printf("%c,%d\n",i,sli[i])
		}
	}

}
//2.利用map
func demo1(s string)map[string]int {
	map1 := make(map[string]int)
	spli := strings.Split(s, "")
	for _,v := range spli {  //v代表的是 字符串中的字符，在map中表示的key
		value,ok := map1[v]//map1的key为每个字符，value为出现的次数
		if ok {
			value++
		} else {
			value = 1
		}
		map1[v]=value
	}
	return map1
}