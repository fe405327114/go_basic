package main

import (
	"os"
	"fmt"
	"io"
)

func main(){
	f,err:=os.OpenFile("D:\\abc.text",os.O_RDWR,6)
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer f.Close()
	n,_:=f.Seek(1,io.SeekStart)
	d:=[]byte("abc")
	f.WriteAt(d,n)
}
