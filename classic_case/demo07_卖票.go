package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

var wg2 sync.WaitGroup
var tickt = 100//票数
func main()  {
	/*
	模拟火车站卖票
	 */
	 wg2.Add(4)
	 //go saleTickets1()
	 //go saleTickets2()
	 //go saleTickets3()
	 //go saleTickets4()
	 go saleTickets("售票口1") //启动子协程，调用saleTickets()
	 go saleTickets("售票口2")//启动字写成，调用saleTickets()
	 go saleTickets("售票口3")
	 go saleTickets("售票口4")

	 wg2.Wait()
	 fmt.Println("程序结束。。")
}

func saleTickets(name string){
	rand.Seed(time.Now().UnixNano())
	for {//g1,g2,g3,g4
		//tickt2 1
		if tickt > 0 {
			time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
			fmt.Println(name,"出售：",tickt) //g3:1,g4:0,g1:-1,g2:-2
			tickt--//0 , -1,-2,-3
		}else {
			fmt.Println(name,"，售票结束")
			break
		}

	}
	wg2.Done()
}
