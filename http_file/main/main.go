//上传（server）
package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

var dir string
var port int

// 初始化参数
func init() {
	dir = path.Dir(os.Args[0])
	flag.IntVar(&port, "port", 9090, "服务器端口")
	flag.Parse()
	fmt.Println("dir:", http.Dir(dir))
}

func main() {
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":"+strconv.Itoa(9090), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Println("ListenAndServer port: ", 9090)

}

// 处理/upload 逻辑
func upload(w http.ResponseWriter, r *http.Request) {
	log.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}

		// fmt.Fprintln("handler.Filename=%s",file.)
		log.Println("handler.Filename=", handler.Filename)
		arr := strings.Split(handler.Filename, "/")
		fileName := arr[len(arr)-1]
		log.Println("after deal, fileName=", fileName)
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		// destFile, err := os.OpenFile("/tmp/upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		destFile, err := os.OpenFile(fmt.Sprintf("/tmp/upload/%s", fileName), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer destFile.Close()
		io.Copy(destFile, file)
		log.Println("upload over...")
	}
}
