package main

import "fmt"

func main(){
	str:="asewgrgngjnfdvnfdhvbbvhjwbv"
	s:=[]byte(str)
	fmt.Printf("%c",s)
	var s1 string
	count:=0
		for i:=0;i<len(s);i++{
			for _,v:=range  s1{
			if string(v)==string(s[i]){
				count++
				break
			}
		}
		if count==0{
			s1+=string (s[i])
		}
	}
	fmt.Println(s1)
}
