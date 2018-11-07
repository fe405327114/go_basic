package main

import "fmt"

//10名学生，3科成绩，打印每个学生成绩，并按平均成绩排名
func main() {
	type student struct {
		id    int
		name  string
		score []float64
		ave   float64
	}
	a1 := []student{student{1, "小明", []float64{13, 24, 25,}, 0},
		student{2, "小红", []float64{13, 24, 25,}, 0},
		student{3, "小花", []float64{13, 24, 25,}, 0},}
		var sum float64
	for i:=0;i<len(a1);i++{
		for j:=0;j<len(a1[i].score);j++{
			sum+=a1[i].score[j]
			a1[i].ave=sum/3
		}
	}
	fmt.Println(a1)
	}
