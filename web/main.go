package main

import (
	"encoding/json"
	"fmt"
	"go_dev/web/exec"
	"log"
	"net/http"
	"strings"
)

type ret_json struct {
	Success bool   `json:"success"`
	Detail  string `json:"detail"`
}

type Profile struct {
	Name    string
	Hobbies []string
}

func checkAndDeal(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/get_tar_file" {
		r.ParseForm() //解析参数，默认是不会解析的
		// fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
		fmt.Println("path", r.URL.Path)
		// fmt.Println("scheme", r.URL.Scheme)
		// fmt.Println("url_long=", r.Form["url_long"])

		user_name := r.Form.Get("user_name")
		pwd := r.Form.Get("pwd")
		wget_url := r.Form.Get("wget_url")
		fmt.Println("user_name=", user_name)
		fmt.Println("pwd=", pwd)
		fmt.Println("wget_url=", wget_url)

		result := ret_json{Success: false, Detail: "deal false."}

		if strings.Trim(user_name, " ") == "test" && strings.Trim(pwd, " ") == "123" && len(wget_url) != 0 {
			// for k, v := range r.Form {
			// 	fmt.Println("key:", k)
			// 	fmt.Println("val:", strings.Join(v, ","))
			// }
			fmt.Println("login success.")
			exitCode := exec.Exec(wget_url)
			if exitCode != 0 {
				result = ret_json{Success: false, Detail: "exec error."}
				json.NewEncoder(w).Encode(result)
			} else {
				// 返回结果
				// fmt.Fprintf(w, "exec success!") //这个写入到w的是输出到客户端的
				result = ret_json{Success: true, Detail: "exec success."}
				json.NewEncoder(w).Encode(result)
			}

		} else {
			fmt.Println("user info error...")

			// 返回结果
			// fmt.Fprintf(w, "param info error...")
			result = ret_json{Success: false, Detail: "param info error..."}
			json.NewEncoder(w).Encode(result)
		}
	} else {
		fmt.Println("path error: ", r.URL.Path)
		json.NewEncoder(w).Encode(ret_json{Success: false, Detail: "path error..."})
	}

}

func ret_json1(w http.ResponseWriter, r *http.Request) {
	result := ret_json{Success: true, Detail: "deal success."}
	fmt.Println(result)

	// b, err := json.Marshal(result)
	// if err != nil {
	// 	log.Println("json format error:", err)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(b)

	// fmt.Println(json.NewEncoder(w).Encode(u))
	json.NewEncoder(w).Encode(result)

}

func foo(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/get_tar_file", checkAndDeal) //设置访问的路由
	http.HandleFunc("/ret_json1", ret_json1)       //设置访问的路由
	http.HandleFunc("/", foo)
	fmt.Println("listen :9090")
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
