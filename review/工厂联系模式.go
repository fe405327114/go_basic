package main

import "fmt"

type Interface interface {
	Jisuan() int
}
type Add struct {
	num1 int
	num2 int
}
type Sub struct {
	Add
}

func (a *Add) Jisuan() (value int) {
	value = a.num1 + a.num2
	return
}
func (s *Sub) Jisuan() (value int) {
	value = s.num1 - s.num2
	return
}

type Factory struct {
}

//func DuoTai(j Interface) (value int) {
//	value=j.Jisuan()
//	return
//}
func (f *Factory) Oper(num1 int, num2 int, OperSign string) (value int) {
	//var j Interface
	switch OperSign {
	case "+":
		Add := Add{num1, num2}
		value=Add.Jisuan()
		//j = &Add
	case "-":
		Sub := Sub{Add{num1, num2}}
		value=Sub.Jisuan()
		//j = &Sub,num2
	}
	//value = DuoTai(j)
	return
}
func main() {
	var f Factory
	value := f.Oper(10, 20, "-")
	fmt.Println(value)

}
