package main

import (
	"encoding/json"
	"fmt"
	"go_dev/http_store_fair_scheduler/exec"
	"log"
	"net/http"
	"strings"
)

type ret_json struct {
	Success bool   `json:"success"`
	Detail  string `json:"detail"`
}

/**
日志输出结构设置
*/
func init() {
	log.SetPrefix("【UserCenter】")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}

func checkAndDeal(w http.ResponseWriter, r *http.Request) {
	if r.Method == "post" {
		r.ParseForm() //解析参数，默认是不会解析的
		log.Println("path", r.URL.Path)
		log.Println(r.Form) //这些信息是输出到服务器端的打印信息

		user_name := r.Form.Get("user_name")
		pwd := r.Form.Get("pwd")
		filePath := r.Form.Get("filePath")
		xmlStr := r.Form.Get("xmlStr")
		log.Println("user_name=", user_name)
		log.Println("pwd=", pwd)
		log.Println("filePath=", filePath)
		log.Println("xmlStr=", xmlStr)

		result := ret_json{Success: false, Detail: "deal false."}

		if strings.Trim(user_name, " ") == "root" && strings.Trim(pwd, " ") == "123" && len(filePath) != 0 && len(xmlStr) != 0 {
			log.Println("login success.")
			// 开始存储文件到本地
			exitCode := exec.StoreFile(filePath, xmlStr)
			if exitCode != 0 {
				result = ret_json{Success: false, Detail: "exec error."}
				json.NewEncoder(w).Encode(result)
			} else {
				// 返回结果
				log.Println("exec success!")
				result = ret_json{Success: true, Detail: "exec success."}
				json.NewEncoder(w).Encode(result)
			}

		} else {
			log.Println("user info error...")

			// 返回结果
			result = ret_json{Success: false, Detail: "param info error..."}
			json.NewEncoder(w).Encode(result)
		}
	} else {
		log.Println("method error: ", r.URL.Path)
		json.NewEncoder(w).Encode(ret_json{Success: false, Detail: "path error..."})
	}

}

func ret_json1(w http.ResponseWriter, r *http.Request) {
	result := ret_json{Success: true, Detail: "deal success."}
	fmt.Println(result)
	json.NewEncoder(w).Encode(result)

}

/**
统一方法入口
*/
func logPanics(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		handle(writer, request)
	}
}

func main() {
	http.HandleFunc("/store_config_file", logPanics(checkAndDeal)) //设置访问的路由
	http.HandleFunc("/ret_json1", logPanics(ret_json1))            //设置访问的路由
	fmt.Println("listen :9090")
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
