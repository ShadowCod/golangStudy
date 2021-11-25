package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 使用os.OpenFile()函数打开文件
	filePath := "f:/file1.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open err:", err)
		return
	}
	defer file.Close()
	// 创建一个写入需要使用的对象
	writer := bufio.NewWriter(file)
	str := "hello,golang"
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 超过缓存大小会自动写入文件，但是没有超过的会存在缓存中，需要使用Flush()将其写入文件
	writer.Flush()
}
