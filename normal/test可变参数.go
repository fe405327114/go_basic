package main

import ("fmt")

func main() {

getSum(3,4,5)

}
func getSum (abc ...int) {
	//sum := 0
	//for i := 0; i < len(abc); i++ {
	//	sum += abc[i]
	//	}
	for i,data :=range abc{
fmt.Println("i的值为，data值为",i,data)

	}

}