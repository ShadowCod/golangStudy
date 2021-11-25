package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Emp 结构体
type Emp struct {
	ID   int
	Name string
}

func testIf(w http.ResponseWriter, r *http.Request) {
	// t := template.Must(template.ParseFiles("if.html"))
	// var emps []*Emp
	// enp1 := &Emp{ID: 1, Name: "zhangshan"}
	// emps = append(emps, enp1)
	// t.Execute(w, emps)
	t := template.Must(template.ParseFiles("if.html"))
	t.Execute(w, "狸猫")
}

// testDrfine 测试模版
func testDefine(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "testDefine")
	t := template.Must(template.ParseFiles("define.html"))
	t.ExecuteTemplate(w, "modle", "")
}

func main() {
	http.HandleFunc("/testIf", testIf)
	http.HandleFunc("/testDefine", testDefine)
	http.ListenAndServe(":8080", nil)
}
