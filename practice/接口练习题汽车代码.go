package main

import "fmt"

//1.定义一个发动机作为接口：启动，加速，停止
type IEngine interface {
	start()
	speedUp()
	stop()
}
type YAMAHA struct {
	name string
}

func (y YAMAHA) start()  {
	fmt.Println(y.name,",启动：60迈")
}
func (y YAMAHA) speedUp() {
	fmt.Println(y.name,",加速：80迈")
}
func (y YAMAHA) stop() {
	fmt.Println(y.name,",停止：0迈")
}

type HONDA struct {
	name string
}
func (h HONDA) start()  {
	fmt.Println(h.name,",启动：40迈")
}
func (h HONDA) speedUp()  {
	fmt.Println(h.name,",加速：120迈")
}
func (h HONDA) stop()  {
	fmt.Println(h.name,",停止：0迈")
}

type Car struct {
	engine IEngine //接口作为字段，需要将该接口的任意实现类对象，赋值
}
func (c Car) testIEngine(){
	c.engine.start()
	c.engine.speedUp()
	c.engine.stop()
}

func main()  {
	/*
	定义一个IEngine接口，3个方法：start(),speedup(),stop()，表示启动，加速，停止
定义两个结构体实现该接口：YAMAHA和HONDA
实现方式分别是：
	YAMAHA
		启动：60，加速80，停止0

	HONDA
		启动：40，加速120，停止0

定义一个Car结构体，将IEngine作为字段，再提供一个car的方法：testIEngine()，用于测试该小汽车的发动机。也就是在testIEngine()中调用start(),speedup(),stop()方法。
	 */
	 //e:=YAMAHA{"雅马哈"}
	 e := HONDA{"本田"}
	 c := Car{e}
	 c.testIEngine()
}
