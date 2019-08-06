package exec

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

/**
日志输出结构设置
*/
func init() {
	log.SetPrefix("【UserCenter】")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
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

func ExecGitCommand(filePath string) (retCode int, retMsg string) {
	// 如果文件夹不存在，则返回错误
	arr := strings.Split(filePath, "/")
	fileName := arr[len(arr)-1]
	dir := strings.Join(arr[:len(arr)-1], "/")
	log.Println("fileName = ", fileName)
	log.Println("dir = ", dir)
	if !Exists(dir) {
		log.Println("dir is not existt ")
		retCode = defaultFailedCode
		retMsg = "dir is not existt "
		return
	}

	//command := fmt.Sprintf(`cd %s && git add %s && git commit -m "add git logic" && git push;`, dir, fileName)
	// 提交本目录下所有文件，防止别人做的更改影响当前配置文件提交。
	command := fmt.Sprintf(`cd %s && git add . && git commit -m "add git logic" && git push;`, dir)

	log.Printf("prepare to start cmd: 【%s】...", command)
	//执行命令
	stdout, stderr, exitCode := RunCommand("/bin/bash", "-c", command)
	log.Println("exec over, exitCode = ", exitCode)

	if exitCode != 0 {
		retMsg = fmt.Sprintf("%s --- %s", stdout, stderr)
		retCode = exitCode
		return
	}

	retMsg = "command exec success..."
	retCode = 0
	return
}

/**
保存文件到挂载的网盘，方便salt模块执行wget命令时获取文件
*/
func StoreFile(filePath, xmlStr string) (retCode int, retMsg string) {

	// 如果文件夹不存在，则返回错误
	arr := strings.Split(filePath, "/")
	fileName := filePath[len(arr)-1]
	dir := strings.Join(arr[:len(arr)-1], "/")
	log.Println("fileName = ", fileName)
	log.Println("dir = ", dir)
	if !Exists(dir) {
		log.Println("dir is not existt ")
		retCode = defaultFailedCode
		retMsg = "dir is not existt "
		return
	}

	file, error := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 755)
	if error != nil {
		log.Println("op file error: ", error)
	}

	writer := bufio.NewWriter(file)
	writer.WriteString(xmlStr)
	writer.Flush()

	retCode = 0
	retMsg = "store file success. "
	return
}
