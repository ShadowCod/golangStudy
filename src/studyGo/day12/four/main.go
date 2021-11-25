package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 拷贝文件（图片|视频|音乐）
// 文件拷贝函数
func Copy(srcFileName, dstFileName string) (written int64, err error) {
	// 打开源文件句柄
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open srcFile err:", err)
		return
	}
	defer srcFile.Close()
	// 获取reader操作对象
	reader := bufio.NewReader(srcFile)
	// 打开目标文件句柄
	dstFile, err := os.Open(dstFileName)
	if err != nil {
		fmt.Println("open dstFile err:", err)
		return
	}
	defer dstFile.Close()
	// 获取writer操作对象
	writer := bufio.NewWriter(dstFile)
	return io.Copy(writer, reader)
}

// 创建一个结构体统计
type CharCount struct {
	ChCount    int //字母的个数
	NumCount   int //数字的个数
	SpaceCount int //空格的个数
	OtherCount int //其他字符的个数
}

// 统计一个文件中字母、数字、空格和其他字符的数量
func Count(path string) {
	// 打开文件句柄
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open file err", err)
	}
	defer file.Close()
	// 获取reader操作对象
	reader := bufio.NewReader(file)
	// 循环读取文件
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读取完毕...")
			break
		}
		// 遍历读取出来的数据进行统计判断
		runeStr := []rune(str)
		var c CharCount
		for _, v := range runeStr {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				c.ChCount++
			case v >= '0' && v <= '9':
				c.NumCount++
			case v == ' ' || v == '\t':
				c.SpaceCount++
			default:
				c.OtherCount++
			}
		}
	}
}
func main() {
	// 命令行参数
	fmt.Println("命令行参数：", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("%v,%v", i, v)
	}
}
