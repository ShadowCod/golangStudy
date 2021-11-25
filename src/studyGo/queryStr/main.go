package main

import (
	"fmt"
	"net/http"
)

/*
	使用golang自带的获取get参数的方法
*/
func getParam(w http.ResponseWriter,r *http.Request){
//	使用FromValue方法取值
	value:=r.FormValue("name")
	fmt.Printf("value:%T",value)
	fmt.Fprintln(w,value)
}
func main(){
	http.HandleFunc("/getParam",getParam)
	err :=http.ListenAndServe(":9090",nil)
	if err !=nil{
		fmt.Printf("listen fail,err:%v\n",err)
		return
	}
}