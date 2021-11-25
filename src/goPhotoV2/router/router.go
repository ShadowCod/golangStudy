package router

import (
	"goPhotoV2/control"
	"log"
	"net/http"
)

// Run 路由
func Run() {
	http.HandleFunc("/index", control.IndexView)
	http.HandleFunc("/load", control.LoadUp)
	http.HandleFunc("/upload", control.APIUpload)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/show", control.ShowPhoto)
	log.Println("Server Run...")
	http.ListenAndServe(":8080", nil)
}
