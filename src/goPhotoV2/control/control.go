package control

import (
	"bytes"
	"goPhotoV2/model"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// IndexView 页面
func IndexView(w http.ResponseWriter, r *http.Request) {
	html := loadHTML("./view/index.html")
	w.Write(html)
}

// LoadUp 上传
func LoadUp(w http.ResponseWriter, r *http.Request) {
	html := loadHTML("./view/load.html")
	w.Write(html)
}

// APIUpload 接收上传的图片保存到数据库
func APIUpload(w http.ResponseWriter, r *http.Request) {
	f, h, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte("<html><head></head><body><h3>FormFileError</h3></body></html>"))
		return
	}
	defer f.Close()
	t := h.Header.Get("Content-Type")
	if !strings.Contains(t, "image") {
		io.WriteString(w, "file type err")
		return
	}
	// err = os.Mkdir("./static", 0666)
	// if err != nil {
	// 	io.WriteString(w, "create dir err")
	// 	return
	// }
	name := time.Now().Format("20060102150405") + path.Ext(h.Filename)
	file, err := os.Create("./static/" + name)
	if err != nil {
		io.WriteString(w, "create file err")
		return
	}
	defer file.Close()
	io.Copy(file, f)
	note := r.FormValue("note")
	mod := &model.Info{
		Name: h.Filename,
		Path: "/static/" + name,
		Unix: time.Now().Unix(),
		Note: note,
	}
	err = model.AddInfo(mod)
	if err != nil {
		io.WriteString(w, "保存sql信息失败")
		return
	}
	io.WriteString(w, "success upload")
}

// ShowPhoto 查看图片
func ShowPhoto(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		io.WriteString(w, "转换string错误")
		return
	}
	mod, err := model.GetInfoByID(num)
	if err != nil {
		io.WriteString(w, "获取数据库信息错误")
		return
	}
	html := loadHTML("./view/detail.html")
	html = bytes.Replace(html, []byte("@src"), []byte(mod.Path), 1)
	w.Write(html)
}

// loadHTML 将html页面转换为[]byte
func loadHTML(name string) []byte {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		return []byte("<html><head></head><body><h3>ReadFileError</h3></body></html>")
	}
	return buf
}
