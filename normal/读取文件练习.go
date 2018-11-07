package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
)

//块读取和换行读取
func main01() {
	f, e := os.OpenFile("D:\\abc.text", os.O_RDWR, 6)
	if e != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer f.Close()
	b := make([]byte, 20)
	for {
		n, e := f.Read(b)
		if e == io.EOF {
			break
		}
		fmt.Print(string(b[:n])) //此处如果用换行打印，就会出现乱码
	}
}
func main() {
	f, err := os.Open("D:\\abc.text")
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer f.Close()
	r:=bufio.NewReader(f)
	for{
		n,e:=r.ReadBytes('\n')//这里要用单引号，不是双引号
		fmt.Print(string(n))
		if e==io.EOF{  //如果将判断放在打印之前，将会缺失最后一行，因为最后一行没有\n
			break
		}
	}
}
