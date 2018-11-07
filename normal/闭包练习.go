package main

import "fmt"

func name() func()int{
	var a int
	return func() int{
		a++
		return a
	}
}

func main (){
r:= name()
for i:=1;i<10;i++{
	fmt.Println(r())
}

}