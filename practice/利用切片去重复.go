package main

import (
	"fmt"
)


func main() {
	str:="sadsdfdsgf"
	var s = []byte(str)  //利用函数传参，将字符串传给切片，否则用range遍历字符串赋值给切片也可以
	var s1 string
	for i := 0; i < len(s); i++ {
		count := 0
		for _,v:=range s1{   //此处必须用range 遍历，因为string类型没有下标，用v来表示s1[i]
			if string(v) == string(s[i]) {
				count++
				break
			}
		}
		if count == 0 {
			s1 += string(s[i])
		}
	}


	fmt.Println(s1)

}

