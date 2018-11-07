package main

import "fmt"

//10名学生，3门成绩，根据平均成绩排序
//一个小组中有3个学生，每个学生有3门课程的成绩需要统计。
//案例要求通过编程依次输入学生的学号、姓名和三门课程的成绩，
//计算出平均成绩并依次把每个学生的学号、姓名和平均成绩打印在屏幕上。
//思路：利用结构体存储三个学生信息，利用切片或者数组存储三科成绩
type studenta struct {
	id    int
	name  string
	score [3]float64
	avg   float64
}

func Avg(s1 [3]studenta) [3]studenta {
	var sum float64
	for i := 0; i < 3; i++ {
		sum = s1[i].score[0] + s1[i].score[1] + s1[i].score[2]
		s1[i].avg = sum / 3
	}
	return s1
}
func main() {
	s1:=[3]studenta{}
	var s2 [3]studenta
	for i := 0; i < len(s1); i++ {
		fmt.Printf("请输出第%d位学生的编号，姓名以及三科成绩\n", i+1)
		fmt.Scan(&s1[i].id, &s1[i].name, &s1[i].score[0], &s1[i].score[1], &s1[i].score[2])
		s2= Avg(s1)
		fmt.Println(s2)
	}
}
