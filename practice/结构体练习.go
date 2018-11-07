package main

import "fmt"

//1创建一个结构体Employee,设置字段id,name,age,salary，创建几个结构体对象，分别赋值，并打印。
//2创建一个结构体User，设置字段username，password，创建几个结构体对象，分别赋值，并打印。
//3.定义一个结构体表示车轮：内半径，外半径。再定义一个结构体表示小汽车：车名，4个车轮。
//4创建一个车结构体，提供属性：颜色，速度。方法：移动()。停止()。创建两个结构体：自行车，跑车，继承车的结构体。分别新增属性和方法。创建对象，进行测试。
//5创建Person结构体。属性：姓名，年龄。方法：显示信息。创建两个结构体：学生和工人。继承Person。分别新增属性成绩和工资
type Employee struct {
	id   int
	name string
	age int
	salary string
}
func main(){
	var s Employee
	s.age=18
	s.name="zhangsan"
	s.id=345
	s.salary="ahahah"
	fmt.Println("s=",s)

	s1:=Employee{id:18, name:"lisi", age:365, salary:"ahaa",}
	fmt.Println("s1=",s1)
	s2:=Employee{13,"tht",43,"rty"}
	fmt.Println("s2=",s2)
	s4:=new(Employee)
	s4.name="dfdg"
	s4.age=45
	fmt.Println("s4=",s4)

	type gongren struct {
		Employee
	}
	s5:=gongren{Employee{13,"ht",34,"rt"}}
	s5.age=13
	fmt.Println("s5=",s5)
	}

