package main

import (
	"fmt"
	"os"
)

func main() {
	// // 将一个文件中的内容导入到另一个文件中
	// filePath1 := "f:/file.txt"
	// filePath2 := "f:/file1.txt"
	// // 读取文件中的内容
	// str, err := ioutil.ReadFile(filePath1)
	// if err != nil {
	// 	fmt.Println("read file err:", err)
	// 	return
	// }
	// // 将读取到的内容写入文件
	// err = ioutil.WriteFile(filePath2, str, 7777)
	// if err != nil {
	// 	fmt.Println("write file err:", err)
	// }
	res, err := PathExists("f:/file.txt")
	fmt.Println(res, err)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
