package main

import "fmt"

//定义一个IEngine接口，3个方法：start(),speedup(),stop()，表示启动，加速，停止
//定义两个结构体实现该接口：YAMAHA和HONDA
//实现方式分别是：
//YAMAHA
//启动：60，加速80，停止0
//
//HONDA
//启动：40，加速120，停止0
//定义一个Car结构体，将IEngine作为字段，再提供一个car的方法：testIEngine()，
// 用于测试该小汽车的发动机。也就是在testIEngine()中调用start(),speedup(),stop()方法。
type IEngine interface {
	start()
	speedup()
	stop()
}

type YAMAHA struct {
	name string
}
type HONGDA struct {
	name string
}

func (y YAMAHA) start() {
		fmt.Println("现在",y.name,"开始启动40迈")
}
func (y YAMAHA) speedup() {
		fmt.Println("现在",y.name,"开始加速80迈")
}
func (y YAMAHA) stop() {
		fmt.Println("现在",y.name,"停止0迈")
}
func (h HONGDA) start() {
	fmt.Println("现在",h.name,"开始启动40迈")
}
func (h HONGDA) speedup() {
	fmt.Println("现在",h.name,"开始加速120迈")
}
func (h HONGDA) stop() {
	fmt.Println("现在",h.name,"停止0迈")
}

type car struct {
	IEngine IEngine
}

func (c car) testIEngine() {
c.IEngine.start()
	c.IEngine.speedup()

	c.IEngine.stop()
}
func main() {
	y1 := YAMAHA{"超音速"}
	h1:=HONGDA{"本田"}// 以结构体.方法来调用，所以先定义结构体
c1:=car{y1}//通过往car结构体里面放置y1  来调用
c2:=car{h1}
	c1.testIEngine()// 以结构体.方法来调用
c2.testIEngine()

}
