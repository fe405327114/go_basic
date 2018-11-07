package main

import (
	"os"
	"strings"
	"fmt"
	"io"
)

// 统计指定目录内单词出现次数：
//（统计指定目录下，所有.txt文件中，“Love”这个单词 出现的次数。）

func main() {
	path := "E:/学习资料"
	//打开目录
	dirInfo, errDir := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if errDir != nil {
		panic(errDir)
	}
	//读取文件
	fileInfos, errFile := dirInfo.Readdir(-1)
	if errFile != nil {
		panic(errFile)
	}
	//判断文件后缀名
	//chan:=make(chan ,)
	for _, fileInfo := range fileInfos {
		if strings.HasSuffix(fileInfo.Name(), ".txt") {
			//如果是txt文件就读取
			hindleTXT(fileInfo)
		}
	}
}
func hindleTXT(fileInfo os.FileInfo) {
	//拼接路径
	filePath := "E:/学习资料/" + fileInfo.Name()
	//打开文件
	f, errOpen := os.OpenFile(filePath, os.O_RDWR, 6)
	if errOpen != nil {
		panic(errOpen)
	}
	//读取文件内容
	result := ""
	buf := make([]byte, 4096)
	for {
		n, errRead := f.Read(buf)
		if errRead != nil {
			//判断是否读完
			if errRead == io.EOF {
				fmt.Println(fileInfo.Name(), "读取完毕")
				break
			} else {
				panic(errRead)
			}
		}
		result += string(buf[:n])
	}
	//统计result中love出现的次数
	//fmt.Println(result)
	count:=strings.Count(result,"love")
	fmt.Println("love的个数为：",count)
}

