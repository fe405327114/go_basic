package main

import "fmt"

func main0102(){
	var name,sex string

	var age int
	fmt.Println("输入年龄")
	fmt.Scanf("%d", &age)
	fmt.Println("请输入你的名字")

	fmt.Scan(&name)//如何同时输入两个信息
	fmt.Println("haha")
	fmt.Scanf("%s\n",&sex)
	fmt.Println("姓名:",name,"性别:",sex)//怎么匹配输出
     fmt.Printf("姓名:%s\n性别:%s\n年龄:%d\n",name,sex,age)

}
//%s 字符串  %d 整型  %c 字符 类型 byte    %f 浮点型  %t 布尔型
