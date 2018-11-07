package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	s1 := make([]int, 10)

	for i := 0; i < 10; i++ {
		var x int
		for {
			count := 0
			x = rand.Intn(10)
			for j:= 0; j < i; j++ {
				if x == s1[j] {
					count++
					break
				}
			}
			if count == 0 {
				break
			}
		}
		s1[i] = x
	}
	fmt.Println("未排序前：", s1)
	for i := 0; i < len(s1)-1; i++ {
		for j := 0; j < len(s1)-i-1; j++ {
			if s1[j] < s1[j+1] {
				s1[j], s1[j+1] = s1[j+1], s1[j]
			}
		}
	}
	fmt.Println("排序后：", s1)
}
