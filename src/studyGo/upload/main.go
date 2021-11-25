package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

/*
	go原生框架上传文件
*/
func Index(w http.ResponseWriter,r *http.Request){
//	解析模版
	t:=template.Must(template.ParseFiles("./index.html"))
//	渲染模版
	t.Execute(w,nil)
}

func UpLoad(w http.ResponseWriter,r *http.Request){
//	注意：要使用MultipartForm，必须先使用ParseMultipartForm解析
//	err:=r.ParseMultipartForm(6)
//	if err !=nil{
//		fmt.Printf("parse form fail,err:%v\n",err)
//		return
//	}
//	根据传入参数key获取值
//	f:=r.MultipartForm.File["f"]
//	fmt.Println("f value:",f)
//	将接收到的文件保存到本地（服务器）

	r.ParseMultipartForm(32<<20)
	f,h,err:=r.FormFile("f")
	if err != nil{
		fmt.Println("form file fail,err:",err)
		return
	}
	defer f.Close()//使用完后应该关闭
	//fmt.Println("f:---",f)
//	创建保存文件
	des,err:=os.Create("./"+h.Filename)
	if err != nil{
		fmt.Println("file create fail,err:",err)
		return
	}
	defer des.Close()
//	读取表单文件，写入保存文件
	_,err=io.Copy(des,f)
	if err != nil{
		fmt.Println("file write fail,err:",err)
		return
	}
}
func main(){
	//访问页面
	http.HandleFunc("/index",Index)
	//指定处理上传文件请求的函数
	http.HandleFunc("/upload",UpLoad)
	//	监听端口
	http.ListenAndServe(":9090",nil)
}
