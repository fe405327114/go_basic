package main

import "fmt"

//通过键盘输入一个字符数组 统计字符出现的次数
func main(){
	var q [20]byte
    fmt.Println("请输入20位字符")

	for i:=0;i<len(q);i++{
		fmt.Scanf("%c",&q[i])
	}

	for i:=0;i<len(q) ;i++{
		fmt.Printf("%c",q[i])

	}
	fmt.Println()
	var as byte
	count :=0
	fmt.Println("请输入查找字符")
	fmt.Scanf("%c\n",&as)
	for i:=0;i<len(q);i++{
		if as==q[i] {
			count++
		}
	}
	fmt.Println(count)
}

