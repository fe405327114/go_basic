package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
)

func main05(){
	f,e:=os.OpenFile("D:/abc.text",os.O_RDWR,6)
	if e!=nil{
		fmt.Println(e)
		return
	}
	defer f.Close()
	b:=make([]byte,20)
	for {
		n,e:=f.Read(b)
		if e == io.EOF {
			break
		}
		fmt.Print(string(b[:n]))
	}
}
func main06(){
	f,e:=os.OpenFile("D:/abc.text",os.O_RDWR,6)
	if e!=nil{
		fmt.Println(e)
		return
	}
	defer f.Close()
	r:=bufio.NewReader(f)
	for   {
		n,e:=r.ReadBytes('\n')
		if e==io.EOF{
			break
		}
		fmt.Print(string((n)))
	}
}
func main(){
	f1,e1:=os.OpenFile("D:/abc.text",os.O_RDWR,6)
	f2,e2:=os.Create("D:/qwe.text")
	if e1!=nil || e2!=nil{
		fmt.Println("-----")
	}
	defer f1.Close()
	defer f2.Close()
	g:=[]byte("白雪公主与大鱼鱼")
	f2.Write(g)

	b:=make([]byte,20)
	for {
		n, e := f1.Read(b)
		if e == io.EOF {
			break
		}
		f2.Write(b[:n])
	}
}