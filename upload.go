package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w, "<html><head><title>上传</title></head>"+
			"<body><form action='#' method=\"post\" enctype=\"multipart/form-data\">"+
			"<label>上传</label>"+":"+
			"<input type=\"file\" name='file'  /><br/><br/>    "+
			"<label><input type=\"submit\" value=\"上传\"/></label></form></body></html>")
	} else {
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		//创建文件
		fW, err := os.Create("/home/py/" + head.Filename)
		if err != nil {
			fmt.Println("文件创建失败")
			return
		}
		defer fW.Close()
		_, err = io.Copy(fW, file)
		if err != nil {
			fmt.Println("文件保存失败")
			return
		}
		//io.WriteString(w, head.Filename+" 保存成功")
		http.Redirect(w, r, "/success", http.StatusFound)
	}

}

func success(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "上传成功")
}

func main() {
	http.HandleFunc("/", upload)
	http.HandleFunc("/success", success)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
