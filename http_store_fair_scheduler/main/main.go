package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"go_dev/http_store_fair_scheduler/exec"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	//获取文件内容
	GETFILECONTENT = 2
	//获取文件
	GETFILE    = 1
	uploadPath = "/data1/upload"
)

type ret_json struct {
	Success bool   `json:"success"`
	Detail  string `json:"detail"`
}

type ret_file_content_json struct {
	Success bool   `json:"success"`
	Data    string `json:"file_data"`
	Detail  string `json:"detail"`
}

/**
日志输出结构设置
*/
func init() {

	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Ldate)
}

/**
保存文件到挂载的网盘，方便salt模块执行wget命令时获取文件
*/
func storeConfigFile(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("【StoreFile】")
	if r.Method == "POST" {
		r.ParseForm() //解析参数，默认是不会解析的
		log.Println("path", r.URL.Path)
		//log.Println(r.Form) //这些信息是输出到服务器端的打印信息

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
			exitCode, retMsg := exec.StoreFile(filePath, xmlStr)
			if exitCode != 0 {
				result = ret_json{Success: false, Detail: "exec error: " + retMsg}
				json.NewEncoder(w).Encode(result)
			} else {
				// 返回结果
				log.Println("exec success!")
				result = ret_json{Success: true, Detail: "exec success."}
				json.NewEncoder(w).Encode(result)
			}

		} else {
			log.Println("param info error...")

			// 返回结果
			result = ret_json{Success: false, Detail: "param info error..."}
			json.NewEncoder(w).Encode(result)
		}
	} else {
		log.Println("method error: ", r.URL.Path)
		json.NewEncoder(w).Encode(ret_json{Success: false, Detail: "path error..."})
	}

}

/**
获取配置文件或文件内容
*/
func getConfigFile(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("【getConfigFile】")
	if r.Method == "POST" {
		r.ParseForm() //解析参数，默认是不会解析的
		log.Println("path", r.URL.Path)

		user_name := r.Form.Get("user_name")
		pwd := r.Form.Get("pwd")
		filePath := r.Form.Get("filePath")
		//获取数据类型：1-获取文件；2-获取文件内容
		getType := r.Form.Get("getType")
		log.Println("user_name=", user_name)
		log.Println("pwd=", pwd)
		log.Println("filePath=", filePath)
		log.Println("getType=", getType)

		result := ret_file_content_json{Success: false, Detail: "deal false."}

		if strings.Trim(user_name, " ") == "root" && strings.Trim(pwd, " ") == "123" && len(filePath) != 0 && len(getType) != 0 {
			log.Println("login success.")

			/* 获取文件内容*/
			exitCode, retMsg := exec.GetFileContent(filePath)
			if exitCode != 0 {
				result = ret_file_content_json{Success: false, Detail: "exec error: " + retMsg}
				json.NewEncoder(w).Encode(result)
			} else {
				// 返回结果
				log.Println("exec success!")
				result = ret_file_content_json{Success: true, Data: retMsg, Detail: "exec success."}
				json.NewEncoder(w).Encode(result)
			}

		} else {
			log.Println("param info error...")

			// 返回结果
			result = ret_file_content_json{Success: false, Detail: "param info error..."}
			json.NewEncoder(w).Encode(result)
		}
	} else {
		log.Println("method error: ", r.URL.Path)
		json.NewEncoder(w).Encode(ret_json{Success: false, Detail: "path error..."})
	}

}

/**
执行git命令，提交网盘
*/
func execGitCommand(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("【execGitCommand】")
	if r.Method == "POST" {
		r.ParseForm() //解析参数，默认是不会解析的
		log.Println("path", r.URL.Path)
		log.Println(r.Form) //这些信息是输出到服务器端的打印信息

		user_name := r.Form.Get("user_name")
		pwd := r.Form.Get("pwd")
		filePath := r.Form.Get("filePath")
		log.Println("user_name=", user_name)
		log.Println("pwd=", pwd)
		log.Println("filePath=", filePath)

		result := ret_json{Success: false, Detail: "deal false."}

		if strings.Trim(user_name, " ") == "root" && strings.Trim(pwd, " ") == "123" && len(filePath) != 0 {
			log.Println("login success.")
			// 开始执行git命令
			exitCode, retMsg := exec.ExecGitCommand(filePath)
			if exitCode != 0 {
				result = ret_json{Success: false, Detail: "exec error: " + retMsg}
				json.NewEncoder(w).Encode(result)
			} else {
				// 返回结果
				log.Println("exec success!")
				result = ret_json{Success: true, Detail: "exec success."}
				json.NewEncoder(w).Encode(result)
			}

		} else {
			log.Println("param info error...")

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
统一方法入口，方便异常处理
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

// 处理/upload 逻辑
func upload(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("【upload】")
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
		srcFile, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}

		// fmt.Fprintln("handler.Filename=%s",srcFile.)
		log.Println("handler.Filename=", handler.Filename)
		arr := strings.Split(handler.Filename, "/")
		fileName := arr[len(arr)-1]
		log.Println("after deal, fileName=", fileName)
		defer srcFile.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		// destFile, err := os.OpenFile("/tmp/upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		destFile, err := os.OpenFile(fmt.Sprintf("%s/%s", uploadPath, fileName), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer destFile.Close()
		io.Copy(destFile, srcFile)
		log.Println("upload over...")
	}
}

func main() {
	// 需要两台机器做负载
	http.HandleFunc("/store_config_file", logPanics(storeConfigFile)) //保存配置文件
	http.HandleFunc("/exec_git_command", logPanics(execGitCommand))   //git命令执行，args：目录、命令
	http.HandleFunc("/get_config_file", logPanics(getConfigFile))     //获取配置文件或文件内容
	http.HandleFunc("/ret_json1", logPanics(ret_json1))               //设置访问的路由
	// 上传
	http.HandleFunc("/upload", logPanics(upload))
	// 下载
	fs := http.FileServer(http.Dir(uploadPath))
	http.Handle("/files/", http.StripPrefix("/files", fs))

	fmt.Println("listen :9090")

	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	/*xmlStr := "bbb"
	//exitCode, retMsg := exec.StoreFile("/tmp/tw.txt", xmlStr)
	err := ioutil.WriteFile("/tmp/tw.txt", []byte(xmlStr), 0755)
	if err != nil {
		fmt.Println(err)
	}*/

	//filePath := "/data11/git/hadoop/druid/hadoop/etc/public/whitelist.xml"
	//filePath := "/data11/git/hadoop/druid/hadoop/etc/public/fair-scheduler.xml"
	//arr := strings.Split(filePath, "/")
	//fileName := arr[len(arr)-1]
	//fmt.Println("file-name = ", fileName)

}
