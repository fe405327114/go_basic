package main

import "fmt"

//"ECALIYHWEQAEFSZCVZTWEYXCPIURVCSWTDBCIOYXGTEGDTUMJHUMBJKHFGUKNKN"

func main01() {
	str := "ECALIYHWEQAEFSZCVZTWEYXCPIURVCSWTDBCIOYXGTEGDTUMJHUMBJKHFGUKNKN"
	arr := [123]int{}
	for i := 0; i < len(str); i++ {
		arr[str[i]]++
	}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			fmt.Printf("%c,%d\n", i, arr[i])
		}
	}
}

func main() {
	str := "ECALIYHWEQAEFSZCVZTWEYXCPIURVCSWTDBCIOYXGTEGDTUMJHUMBJKHFGUKNKN"
	arr := []byte(str)
	m1 := make(map[string]int)
	for i := 0; i < len(arr); i++ {
		value, ok := m1[string(arr[i])]
		if ok {
			value++
		} else {
			value=1
		}
		m1[string(arr[i])]=value
	}
	fmt.Println(m1)
}
