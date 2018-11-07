package main

import "fmt"

//一个小组中有3个学生，每个学生有3门课程的成绩需要统计。
//案例要求通过编程依次输入学生的学号、姓名和三门课程的成绩，
//计算出平均成绩并依次把每个学生的学号、姓名和平均成绩打印在屏幕上。
//思路：利用结构体存储三个学生信息，利用切片或者数组存储三科成绩
type student struct {
	id int
	name string
	score [3]float64
	avge float64
}
func Avg(stu [3]student)[3]student{
	var sum float64
	for i:=0;i<len(stu);i++{
		sum=stu[i].score[0]+stu[i].score[1]+stu[i].score[2]
		stu[i].avge=sum/3
	}
	return  stu
}
func main(){
	students:=[3]student{}
	var stus [3]student
	for i:=0;i<3;i++{
		fmt.Printf("请输入第%d个学生的编号，姓名，成绩1，成绩2，成绩3\n",i+1)
		fmt.Scan(&students[i].id,&students[i].name,&students[i].score[0],
			&students[i].score[1],&students[i].score[2])
		stus=Avg(students)
		fmt.Println(stus)
	}
}

