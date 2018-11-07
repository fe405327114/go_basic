package main

import "fmt"

type Oper interface {
	Jisuan() int
}

type Add struct {
	num1 int
	num2 int
}

type Sub struct {
	Add
}

func (a Add) Jisuan() (value int) {
	value = a.num1 + a.num2
	return
}
func (s Sub) Jisuan() (value int) {
	value = s.num1 + s.num2
	return
}
func Duotai(o Oper) (value int) {
	value = o.Jisuan()
	return
}

type Factory struct {
}

func (f Factory) Cal(num1 int, num2 int, Sign string) (value int) {
	var o Oper
	switch Sign {
	case "+":
		Add := Add{num1,num2}
		//value=Add.Jisaun()
		o = &Add
	case "-":
		Sub := Sub{Add{num1, num2}}
		//value=Sub.Jisaun()
		o = &Sub
	}
	value = Duotai(o)
	return
}
func main(){
	var f Factory
	value:=f.Cal(10,20,"+")
	fmt.Println(value)
}
