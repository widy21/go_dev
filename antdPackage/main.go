package main

import (
	"fmt"
	"go_dev/http_store_fair_scheduler/exec"
	"log"
	"time"
)

/**
日志输出结构设置
*/
func init() {

	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Ldate)
}

/**
执行git命令，提交网盘
*/
func execCommand() {
	log.SetPrefix("【execCommand】")

	src_dir := "/Users/huaxiao/Documents/workspace/java/jmr-portal-antdesignpro"
	target_dir := "/Users/huaxiao/Documents/workspace/java/jmr-portal-service/src/main/resources/static"
	timeStr := time.Now().Format("200601021504")

	//执行命令1
	command := fmt.Sprintf(`cd %s && git add . && git commit -m 'fix bug %s' --no-verify && git pull && git push && rm -rf ./dist && npm run-script build;`, src_dir, timeStr)
	//command := fmt.Sprintf(`cd %s && rm -rf ./dist && npm run-script build;`, src_dir)
	log.Printf("prepare to start cmd: 【%s】...", command)
	stdout, stderr, exitCode := exec.RunCommand("/bin/bash", "-c", command)
	log.Println("exec command1 over, exitCode = ", exitCode)
	if exitCode != 0 {
		retMsg := fmt.Sprintf("%s --- %s", stdout, stderr)
		log.Panicln(retMsg)
	}

	//判断文件夹是否存在
	for i := 0; i < 10; i++ {
		log.Println("check exist src_dir for", i, "time.")
		if !exec.Exists(src_dir) {
			time.Sleep(10 * time.Second)
		} else {
			break
		}
	}

	//执行命令2
	command = fmt.Sprintf(`cd %s && rm -rf ./* && cp -rf %s/dist/* %s/ && git add . && git commit -m 'fix bug %s' && git pull && git push`, target_dir, src_dir, target_dir, timeStr)
	stdout, stderr, exitCode = exec.RunCommand("/bin/bash", "-c", command)
	log.Println("exec command2 over, exitCode = ", exitCode)
	if exitCode != 0 {
		retMsg := fmt.Sprintf("%s --- %s", stdout, stderr)
		log.Panicln(retMsg)
	}
}

func main() {
	//execCommand()
	//timeStr := time.Now().Format("20060102150405")
	timeStr := time.Now().Format("200601021504")
	log.Println(timeStr)
}
