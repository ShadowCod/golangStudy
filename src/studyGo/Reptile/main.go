package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
)

// 定义正则匹配
var (
	reImg      = ``
	chSem      = make(chan int, 5) // 设置限制协程数量的管道(控制信号量)
	downloadGW sync.WaitGroup      //等待组
)

// GetHTML 根据传入的url返回网页的html
func GetHTML(url string) (html string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get err", err)
		return
	}
	// 因为response.Body是一个io,不使用时需要关闭
	defer response.Body.Close()
	// 获取body中的信息
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll err", err)
		return
	}
	// 将[]byte转换为string类型,转换完成后就是网页代码
	html = string(bytes)
	return
}

// GetUtlSlice 获取切片
func GetUtlSlice() []string {
	html := GetHTML("http://www.baidu.com")
	fmt.Println(html)
	// 生成正则对象
	re := regexp.MustCompile(reImg)
	// 找出所有符合正则的,得到的结果是一个二维切片
	strL := re.FindAllStringSubmatch(html, -1)
	imgURLS := make([]string, 0)
	// fmt.Println(strL)
	for _, v := range strL {
		// fmt.Println(v[1])
		imgURLS = append(imgURLS, v[1])
	}
	return imgURLS
}

// DownLoadImg 下载图片的函数(同步)
func DownLoadImg(s []string, filename string) {
	for _, v := range s {
		// 获取到的resp
		resp, _ := http.Get(v)
		img, _ := ioutil.ReadAll(resp.Body)
		err := ioutil.WriteFile(filename, img, 0644)
		if err != nil {
			fmt.Println(v, "下载失败")
		} else {
			fmt.Println(v, "下载成功")
		}
	}
}

// DownLoadImgAsync 异步下载图片
func DownLoadImgAsync(url string, filename string) {
	// 每开一个协程,就写一个
	downloadGW.Add(1)
	go func() {
		chSem <- 1
		DownFile(url, filename)
		<-chSem
		// 执行完了就done1
		downloadGW.Done()
	}()
}

// DownFile 下载资源
func DownFile(url, filename string) {
	// 根据url获取resp
	resp, _ := http.Get(url)
	// 根据resp获取[]byte
	bytes, _ := ioutil.ReadAll(resp.Body)
	// 写入本地文件
	err := ioutil.WriteFile(filename, bytes, 0644)
	if err != nil {
		fmt.Println(url, "下载失败")
	} else {
		fmt.Println(url, "下载成功")
	}
}

// GetRandomName 随机名称
func GetRandomName() {

}

func main() {
	imgURLS := GetUtlSlice()
	// 下载函数
	DownLoadImg(imgURLS, "")
	// 让主协程等待
	downloadGW.Wait()
}

// GW 使用方式
func GW() {
	var wg sync.WaitGroup
	url := "www.baidu.com"
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			http.Get(u)
		}(url)
	}
	wg.Wait()
}
