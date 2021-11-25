package main

/*
	用原生的
*/
import (
	"fmt"
	"net/http"
	"html/template"
)

type User struct{
	Name string
	gender string
	Age int
}

func init(){
	fmt.Println("init 函数")
}

func init(){
	fmt.Println("init 函数")
}

func Handle(w http.ResponseWriter,r *http.Request){
//	解析模版
	tem,err :=template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("解析模版失败",err)
		return
	}
//	渲染模版
	user :=&User{
		Name:"悟空",
		gender:"男",
		Age:500,
	}

	list:=[]int{1,2,3,4}
	err = tem.Execute(w,map[string]interface{}{
	"user":user,
	"list":list,
	})
	if err != nil{
		fmt.Printf("渲染模版失败,Err:%v",err)
		return
	}
}
func main(){
	http.HandleFunc("/hello",Handle)
	err := http.ListenAndServe(":9090",nil)
	if err != nil{
		fmt.Println("设置监听端口失败：",err)
		return
	}
}