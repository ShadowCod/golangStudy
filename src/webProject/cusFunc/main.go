package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
	自定义模版函数
*/
func func1(w http.ResponseWriter,r *http.Request){
//	自定义函数   注意：自定义函数要么只有一个返回值，要么第二个返回值必须是error类型
	f:=func(name string)(s string,err error){
	return name +"你好",nil
	}
//	定义模版
//	先创建再解析
	t:=template.New("t.tmpl")//生成一个名字为“”的模版	注意：new中的名字应该和解析的文件名称一致
	//	注册函数   必须在解析模版之前
	t.Funcs(template.FuncMap{
		"sayHello":f,
	})
	//	解析模版
	_,err := t.ParseFiles("./t.tmpl")
	//t,err :=template.ParseFiles("./t.tmpl")
	if err != nil{
		fmt.Println("解析模版出错")
		return
	}
//	渲染模版
	t.Execute(w,"jerry")
}

//嵌套模版
func demo1(w http.ResponseWriter,r *http.Request){
//	定义模版
//	解析模版
	t,err:=template.ParseFiles("./m.tmpl","./ul.tmpl")//注意：被嵌套的应该放在后面
	if err != nil{
		fmt.Printf("parse template err,err:%v",err)
		return
	}
//	渲染模版
	t.Execute(w,"tom")
}

//模版继承
func inherit(w http.ResponseWriter,r *http.Request){
//	定义模版
//	t:=template.New("index.tmpl").Delims()
//	解析模版
	t,err:=template.ParseFiles("./template/base.tmpl","./template/index.tmpl")
	if err != nil{
		fmt.Printf("parse template fail,err:%v",err)
		return
	}
//	渲染模版	注意：渲染多个模版需要使用ExecuteTemplate
	t.ExecuteTemplate(w,"index.tmpl","tom")
}

func main(){
//	生成处理器函数
	http.HandleFunc("/",func1)
	http.HandleFunc("/nested",demo1)
	http.HandleFunc("/inherit",inherit)
//	监听端口
	err :=http.ListenAndServe(":9000",nil)
	if err != nil{
		fmt.Println("监听端口失败")
		return
	}
}
