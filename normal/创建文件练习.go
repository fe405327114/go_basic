package main

import (
	"os"
	"fmt"
)

func main(){
	f,e:=os.Create("D:\\abc.text")
	if e!=nil{
		fmt.Println("文件创建失败")
		return //这里需要有个打断函数继续执行的操作，不然会继续往下执行
	}
	 defer f.Close()
	f.WriteString("asffe\r\n")  //windows系统中这里需要用\r\n来换行，
	s:="fesggrg\r\n"
	b:=[]byte(s)
	f.Write(b)
	s1:="你好啊狗狗\r\n"
	c:=[]byte(s1)
	f.Write(c)
	n,_:=f.Write(c)
	fmt.Println(n)


}
