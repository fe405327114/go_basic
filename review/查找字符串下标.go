package main

import "fmt"

func main(){
	str:="sadsfdsfdsfd"
	b:=[]byte(str)
	for i:=0;i<len(b);i++{
		if b[i]=='s'{
			fmt.Println(i+1)
		}
	}
}