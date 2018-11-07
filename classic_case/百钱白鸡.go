package main

import "fmt"

//一只公鸡值五钱，一只母鸡值三钱，三只小鸡值一钱，现在要用百钱买百鸡，请问公鸡、母鸡、小鸡各多少只？
func main(){
	for cock:=1;cock<20;cock++ {
		for hen:=1;hen<33;hen ++ {
			chiken:=100-cock-hen
			if chiken/3+5*cock+3*hen==100 && chiken%3==0{
				fmt.Println(cock,hen,chiken)
			}

		}
	}

}