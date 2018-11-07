package main

import (
	"fmt"
)

type animal struct {
	name string
	age int
}
func (a animal)eat(){
	fmt.Println("阿猫阿狗爱吃鱼")
}
type cat struct {
	animal
	color string
}
func (a cat)eat (){
fmt.Println("小猫说我更爱舔鱼")
}
type dog struct {
	animal	
}
func(b dog)lookdoor(){
	fmt.Println("我会看门")
}

func main(){
	a:=animal{name:"xiaohua",age:18}
	animal.eat(a)
	m:=cat{color:"red"}
	cat.eat(m)

	fmt.Println()

}

