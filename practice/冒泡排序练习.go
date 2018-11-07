package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	var a [10]int
	n :=len(a)
	for i:=0;i<n;i++{
		a[i]=rand.Intn(100)
		fmt.Printf("%d,",a[i])

	}
    for m:=0;m<n-1;m++{
	for z:=0;z<n-m-1;z++{
		if a[z]>a[z+1]{
			a[z],a[z+1]=a[z+1],a[z]
	    }
	  }
	}
	fmt.Printf("%d",a)
}
