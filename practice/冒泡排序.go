package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var a [5]int
	n := len(a)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100)
		fmt.Printf("%d,", a[i])
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Printf("排序后随机数是:%d\n",a)
}
