package main

import (
	"fmt"
	"image/jpeg"
)

//定义结构体存储5名学生 三门成绩 求出每名学生的总成绩和平均成绩
 type stud struct {
 	id int
 	name string
 	score []int
 }
 func main() {
 	a1:=new()
	 a := []stud{stud{101, "小红", []int{13, 14, 15}},
		 stud{121, "小名", []int{13, 1, 15}},
		 stud{131, "小刚", []int{134, 10, 15}}}
sum:=0
	 for i:=0;i<len(a);i++{
				for j:=0;j<len(a[i].score);j++{
	 		sum+=a[i].score[j]
		}
		 fmt.Printf("第%d名学生成绩为：%d，平均成绩为：%d\n",i+1,sum,sum/3)
	 }

 }