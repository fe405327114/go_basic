package main

import "fmt"

func main() {
	const(
		a=iota
		b=iota
		c
	)

	fmt.Println(a,b,c)
	fmt.Printf("a=%d\nb=%d\nc=%d\n",a,b,c)
}
