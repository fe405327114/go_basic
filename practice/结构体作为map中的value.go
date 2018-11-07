package main

import (
	"fmt"
)

type stu struct {
	name string
	age int
	score int
}
func main() {
	map0 := make(map[int][]stu)
	map0[101] = append(map0[101], stu{"曹操", 14, 78}, stu{"张飞", 13, 56})
	map0[102] = append(map0[102], stu{"曹操", 14, 78}, stu{"张飞", 13, 56})
	for k,v:= range map0 {
		for i,data:=range v {
			fmt.Println(k,i,data)
		}
	}
}

