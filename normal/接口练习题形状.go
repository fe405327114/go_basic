package main

import (
	"fmt"
	"math"
)

//定义一个接口：形状
//	定义两个方法：
//		周长()
//		面积()
//定义三个结构体分别实现该接口：
//	三角形：周长()，面积()
//	矩形：周长()，面积()
//	圆形：周长()，面积()
//海伦公式 s=根号下p(p-a)(p-b)(p-c)  p=L/2
//在主函数中：分别创建三种形状的对象，并调用对应的周长和面积的方法
type shape interface {
	length()  float64
	area()  float64

}
type triangle struct {
	a float64
	b float64
	c float64
}
type rectangle struct {
	shapes string
}
type roundness struct {
	shapes string
}

func (t triangle) length()float64{
	return t.a+t.b+t.c
}
func (t triangle) area()float64{
	p:=t.length()/2
	area:=math.Sqrt(p*(p-t.a)*(p-t.b)*(p-t.c))
	return  area
}
func main()  {


	t1:=triangle{a:3,b:4,c:5}
	t1.length()
	t1.area()
	fmt.Println("三角形的周长为",t1.length(),"三角形的面积为",t1.area())



}