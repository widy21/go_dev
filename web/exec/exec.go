package exec

import (
	"bytes"
	"fmt"
	"go_dev/web/md5file"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func get_file(file_path string) string {
	skillfolder := file_path
	var ret_file string
	// 获取第一个文件
	files, _ := ioutil.ReadDir(skillfolder)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			ret_file = file.Name()
			break
		}
	}
	return ret_file
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

const defaultFailedCode = 1

func RunCommand(name string, args ...string) (stdout string, stderr string, exitCode int) {
	log.Println("run command:", name, args)
	var outbuf, errbuf bytes.Buffer
	cmd := exec.Command(name, args...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	stdout = outbuf.String()
	stderr = errbuf.String()

	if err != nil {
		// try to get the exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			// This will happen (in OSX) if `name` is not available in $PATH,
			// in this situation, exit code could not be get, and stderr will be
			// empty string very likely, so we use the default fail code, and format err
			// to string and set to stderr
			log.Printf("Could not get exit code for failed program: %v, %v", name, args)
			exitCode = defaultFailedCode
			if stderr == "" {
				stderr = err.Error()
			}
		}
	} else {
		// success, exitCode should be 0 if go is ok
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}
	fmt.Printf("command result, stdout: %v, stderr: %v, exitCode: %v", stdout, stderr, exitCode)
	return
}

func Exec(wget_url string) int {
	// 当前时间
	// timeStr:=time.Now().Format("2006-01-02 15:04:05")
	//timeStr := time.Now().Format("20060102")
	//file_path := fmt.Sprintf(`/root/root/%ssre`, timeStr)

	//不再按照时间分包，方便再次执行前日失败的任务 -- 2020-09-17 14:47:15修改。
	file_path := `/root/root/sre`
	if !Exists(file_path) {
		err_msg := fmt.Sprintf("file_path[%s] not exists.", file_path)
		log.Println(err_msg)
		// panic_test(err_msg)
		// 创建文件夹
		err := os.MkdirAll(file_path, os.ModePerm)
		if err != nil {
			err_msg = fmt.Sprintf("mkdir failed![%v]\n", err)
			log.Printf("mkdir failed![%v]\n", err_msg)
			panic(err_msg)
		} else {
			log.Printf("mkdir success!\n")
		}
	}
	arr := strings.Split(wget_url, "/")
	ret_file := arr[len(arr)-1]
	command := fmt.Sprintf(`wget -t 1 --timeout=120 -P %s %s -O %s/%s`, file_path, wget_url, file_path, ret_file)
	// command := fmt.Sprintf(`wget -t 1 --timeout=10 -P %s %s -O %s`, file_path, wget_url, ret_file)
	// cmd := exec.Command("/bin/bash", "-c", `wget -P /tmp/testweb http://www.microbrew.org/tools/md5sha1sum/md5sha1sum-0.9.5.tar.gz`)

	log.Println("command = ", command)
	log.Println("prepare to start cmd...")
	//执行命令
	_, _, exitCode := RunCommand("/bin/bash", "-c", command)
	log.Println("exec over, exitCode = ", exitCode)

	if exitCode != 0 {
		return exitCode
	}

	log.Println("begin to genarate md5 file...")
	log.Printf("ret_file = [%s]\n", ret_file)
	md5file.Md5File(file_path, ret_file)
	log.Println("genarate md5 file over...")
	return 0
}
