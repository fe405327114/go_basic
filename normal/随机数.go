package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	a:=rand.Intn(1000)
    fmt.Println(a)

}