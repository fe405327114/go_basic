package main

import (
	"fmt"
)

type Jiekou interface {
	Jisuan() int
}
type Fulei struct {
	num1 int
	num2 int
}
type Jiafa struct {
	Fulei
}
type Jianfa struct {
	Fulei
}

func (a Jiafa) Jisuan() (value int) {
	value = a.num1 + a.num2
	return
}
func (b Jianfa) Jisuan() (value int) {
	value = b.num1 + b.num2
	return
}

type Gochang struct {
}

func Duotai(j Jiekou)(value int){//多态的实现
	value=j.Jisuan() //接口.方法
	return
}
//工厂设计模式
func (c Gochang) Yunzuo(num1 int, num2 int, Yunsuanfu string) (value int) {
	var j Jiekou
	switch Yunsuanfu {
	case "+":
		var Jiafa Jiafa = Jiafa{Fulei{num1, num2}}
		//value = Jiafa.Jisuan()   //使用多态后，不再同伙对象.方法来调用
		j=&Jiafa
	case "-":
		var Jianfa Jiafa = Jiafa{Fulei{num1, num2}}
		//value = Jianfa.Jisuan()
		j=&Jianfa
	}
	value=Duotai(j)
	return
}
func main() {
	var c Gochang
	value:=c.Yunzuo(10,20,"+")
	fmt.Println(value)
}
