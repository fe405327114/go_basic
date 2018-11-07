package main

import (
	"fmt"
)

//定义结构体存储5名学生 三门成绩 求出每名学生的总成绩和平均成绩，并根据平均成绩排名
type stud struct {
	id    int
	name  string
	score []int
	ave   int
}

func main() {
	a := []stud{stud{101, "小红", []int{13, 14, 15},0},
		stud{121, "小名", []int{13, 1, 15}, 0},
		stud{131, "小刚", []int{134, 10, 15},0}}
	sum := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i].score); j++ {
			sum += a[i].score[j]
			a[i].ave=sum/3
		}
		fmt.Printf("第%d名学生成绩为：%d，平均成绩为：%d\n", i+1, sum, sum/3)
	}
	//利用冒泡排序进行排名
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j].ave<a[j+1].ave{
				a[j],a[j+1]=a[j+1],a[j]
			}
		}
		}
		fmt.Printf("第一名为：%v\n第二名为：%v\n第三名为：%v\n",a[0],a[1],a[2])
}
