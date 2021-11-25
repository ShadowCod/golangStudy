package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

/*
	正则表达式
*/

var (
	// <a href="//www.bilibili.com/video/BV11a4y15788" target="_blank"></a>
	reImg = `<a[\s\S]+?href="(//www.bilibili.com[\s\S]+?)"`
)

func main() {
	resp, err := http.Get("https://www.bilibili.com")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	// 将byte转换为string
	html := string(bytes)
	fmt.Println(html)

	re := regexp.MustCompile(reImg)
	rets := re.FindAllStringSubmatch(html, -1)
	fmt.Println("捕获视频数:", len(rets))
	// for _, ret := range rets {
	// 	fmt.Println(ret[1])
	// }
}
