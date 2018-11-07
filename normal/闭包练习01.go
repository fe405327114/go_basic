package main

import "fmt"

func demo()func ()int{
var a int
return func() int {
	a++
	return a
}
}
func main(){
	b:=demo()
	for i:=0;i<10 ;i++  {
		fmt.Println(b())
	}


}