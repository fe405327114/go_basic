package main

import "fmt"

func main() {

	var year int
	fmt.Println("请输入年份")
	fmt.Scan(&year)
	//b:=year%400==0 || year %4==0 && year%100!=0
	b := year%400 == 0 || year%4 == 0 && year%100 != 0
	//等于要用“==”
	fmt.Println(b)
}
