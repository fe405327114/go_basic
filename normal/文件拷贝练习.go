package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	f1, e1 := os.Open("D：/abc.text")
	f2, e2 := os.Create("D:/egf.text")
	if e1 != nil || e2 != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer f1.Close()
	defer f2.Close()
	s := make([]byte, 1024)
	for{  //利用块读取，不用行读取因为通常情况下某个站位符的位置不固定
		n,e3:=f1.Read(s)
		if e3==io.EOF{
			break
		}
		f2.Write(s[:n])
	}
}
