package main

//定义函数，
type name func(int int)int
s1:=name
//定义数组    有下标，数值可以重复，数组的长度固定的
var arr[5] int
arr[1]=1

var arr[5]int{1,2,3,3,4,}
var arr[5]int{3:5}//表示的是 下标3的数值为5
//定义切片
var art[]int
s2:=make([]int,5,10)
s2=append(s2,2)

//创建map

var m map[int]string
dict:=map[int]string

//结构体
var student struct{
	name string
	id int
}
type studenr struct {
	sex string
	age int
}
st:=student{name:zhang,id 18}

//指针
var i int=100
var p *int//指针变量
p=&i
