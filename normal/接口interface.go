package main

import (
	"fmt"
)

//接口是方法的集合，将方法和具体的功能实现 分离开来。在不同的类中进行实现
//就相当于是函数的集合
type USB interface {
	//只有功能的描述：方法的集合，没有字段
	name() string
	plugin()
}
type mouse struct {
	vi string
}
func (m mouse)name()string{
	return m.name()
}
func(m mouse)plugin(){
	fmt.Println("鼠标品牌是" ,m.vi,"现在请插入")
}
type flashdisk struct {
	length int
}
func(f flashdisk)name()string{
	return f.name()
}
func(f flashdisk)plugin(){
	fmt.Println("U盘的长度是",f.length,"可以插入")
}
func main(){
	//错误的写法，不可以直接mouse.vi
	//要先用变量定义为mouse类型才可以开始初始化
	//如m1：=mouse{}   m1.vi=
	m1:=mouse{"双飞燕"}  //结构体初始化
	m1.plugin()  // 注,意方法调用的格式
	f1:=flashdisk{18}
	f1.plugin()
}