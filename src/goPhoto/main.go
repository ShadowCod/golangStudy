package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// 网页版
func indexView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

// loadHTML 加载Html
func loadHTML(name string) (b []byte) {
	f, err := os.Open(name)
	if err != nil {
		return []byte("<html><head></head><body><h3>Error</h3></body></html>")
	}
	b, err = ioutil.ReadAll(f)
	if err != nil {
		return []byte("<html><head></head><body><h3>Error</h3></body></html>")
	}
	return b
}

// 上传图片
func upLoad(w http.ResponseWriter, r *http.Request) {
	// GET
	if r.Method == "GET" {
		html := loadHTML("./view/upload.html")
		w.Write(html)
		return
	}
	// POST
	if r.Method == "POST" {
		f, h, err := r.FormFile("file")
		if err != nil {
			w.Write([]byte("文件上传出现错误：" + err.Error()))
			return
		}
		os.Mkdir("./static", 0666)
		out, err := os.Create("./static/" + h.Filename)
		if err != nil {
			io.WriteString(w, "创建文件失败:"+err.Error())
			return
		}
		_, err = io.Copy(out, f)
		if err != nil {
			io.WriteString(w, "文件保存失败："+err.Error())
			return
		}
		out.Close()
		f.Close()
		// io.WriteString(w, "保存成功:/static"+h.Filename)
		http.Redirect(w, r, "/detail?name="+h.Filename, 302)
	}
}

// ImageView 展示图片
func ImageView(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	f, err := os.Open("./static/" + name)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	defer f.Close()
	w.Header().Set("Content-Type", "image")
	io.Copy(w, f)
}

// DetailView 网页展示图片
func DetailView(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	html := loadHTML("./view/detail.html")
	html = bytes.Replace(html, []byte("@src"), []byte("/image?name="+name), 1)
	w.Write(html)
}

// 展示列表
func listView(w http.ResponseWriter, r *http.Request) {
	html := loadHTML("./view/list.html")
	names, err := ioutil.ReadDir("./static")
	if err != nil {
		w.Write([]byte("errors"))
		return
	}
	temp := ""
	for i := 0; i < len(names); i++ {
		temp += `<li><a href="/detail?name=` + names[i].Name() + `"><img src="/image?name=` + names[i].Name() + `"> alt = "未发现"></a></li>`
	}
	html = bytes.Replace(html, []byte("@html"), []byte(temp), 1)
	w.Write(html)
}

func main() {
	fmt.Println("hello")
	http.HandleFunc("/index", indexView)
	http.HandleFunc("/load", upLoad)
	http.HandleFunc("/image", ImageView)
	http.HandleFunc("/detail", DetailView)
	http.HandleFunc("/list", listView)
	log.Println("run...")
	http.ListenAndServe(":8080", nil)
}
