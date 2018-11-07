package main

import "fmt"

type Humaner interface {
	SayHi()
}
type Personer interface {
	Humaner
	Sing()
}
type student10 struct {
	name string
	score int
}
type singer10 struct {
	student10
	sings string
}
func (s student10)SayHi(){
	fmt.Printf("我叫%s,我的分数是%d\n",s.name,s.score)
}
func (i singer10)Sing(){
	i.SayHi()
	fmt.Printf("我会唱%s",i.sings)
}
func Sayhello(p Personer){
	p.SayHi()
	p.Sing()
}
func main(){
	var p  Personer
	p=&singer10{student10{"张鹏义",59},"两只老虎"}
	p.Sing()
}