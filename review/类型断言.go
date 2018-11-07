package main

import (
	"fmt"
)

func main (){
	a1:=make([]interface{},3)
	a1[0]="白雪你好"
	a1[1]=1213
	a1[2]=13.2
	for _,v:=range a1{
		if data,ok:=v.(int);ok{
			fmt.Println(data)
		}else if data,ok:=v.(string);ok{
			fmt.Println(data)
		}else  if data,ok:=v.(float64);ok{
			fmt.Println(data)
		}
	}
}
