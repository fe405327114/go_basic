package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

var wg3 sync.WaitGroup
var tickt2 = 10 //票数
var matex sync.Mutex //创建一把锁头
func main()  {
	/*
	上锁：
	 */
	 wg3.Add(4)
	 //go saleTickets1()
	 //go saleTickets2()
	 //go saleTickets3()
	 //go saleTickets4()
	 go saleTickets2("售票口1") //启动子协程，调用saleTickets()
	 go saleTickets2("售票口2") //启动字写成，调用saleTickets()
	 go saleTickets2("售票口3")
	 go saleTickets2("售票口4")

	 wg3.Wait()
	 fmt.Println("程序结束。。")
}

func saleTickets2(name string){
	rand.Seed(time.Now().UnixNano())
	for {//g1,g2,g3,g4
		//tickt2 1
		//上锁
		matex.Lock() //g1 g2
		if tickt2 > 0 {
			time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
			fmt.Println(name,"出售：", tickt2) //g3:1,g4:0,g1:-1,g2:-2
			tickt2--                        //0 , -1,-2,-3
		}else {
			fmt.Println(name,"，售票结束")
			matex.Unlock()
			break
		}
		//开锁
		matex.Unlock()

	}
	wg3.Done()
}
