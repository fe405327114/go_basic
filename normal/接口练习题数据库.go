package main

import "fmt"

//定义一个DataBase接口，4个方法：insert(),update(),query(),delete()
//定义两个结构体实现该接口：Mysql和Oracle
//定义一个Project结构体：两个字段：name表示项目名字，DataBase表示项目用的数据库，
// 定义一个方法：testDataBase(),测试DataBase的4个方法。
type DataBase interface {
	insert()
	update()
	query()
	delete()
}
type Mysql struct {
	name string
}
type Oracle struct {
	name string
}
func(m Mysql)insert(){
	fmt.Println("增",m.name)
}
func(m Mysql)update(){
	fmt.Println("改",m.name)
}
func(m Mysql)query(){
	fmt.Println("查",m.name)
}
func(m Mysql)delete(){
	fmt.Println("删",m.name)
}
type Project struct {
	name string
	DataBase DataBase
}
func(p Project)testDataBase(){
	p.DataBase.delete()
	p.DataBase.insert()
	p.DataBase.update()
	p.DataBase.query()
}
func main(){
	m1:=Mysql{"张三"}
	p1:=Project{"图书馆数据库",m1}
	p1.testDataBase()
}