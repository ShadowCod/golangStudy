package main

import (
	"fmt"
	"io/ioutil"
)

// 文件
func main() {
	// // 打开文件Open(路径)
	// // file:file对象|file指针|文件句柄
	// file, err := os.Open("../../../../file.txt")
	// if err != nil {
	// 	fmt.Println("open err:", err)
	// }
	// fmt.Println(file) //file就是一个指针
	// //关闭文件
	// err = file.Close()
	// if err != nil {
	// 	fmt.Println("close err:", err)
	// }

	// file, err := os.Open("../../../../file.txt")
	// if err != nil {
	// 	fmt.Println("open err:", err)
	// }
	// defer file.Close() //要及时关闭file句柄，否则会有内存泄漏

	// // 创建一个*Reader,带缓冲的
	// reader := bufio.NewReader(file)
	// for {
	// 	str, err := reader.ReadString('\n')
	// 	fmt.Print(str)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }
	// fmt.Println("读取文件结束")

	// 一次性读文件，适用于小文件
	file := "../../../../file.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
}
