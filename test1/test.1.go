package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	i := 100
	var arr = [...]int{1, 2, 3, 5, 01}
	sort.Ints(arr[:])
	fmt.Println("hello world~", i, arr)

	// 文件路径截取
	wgetUrl := "/data11/git/hadoop/evil/yarn/etc/public_new/fair-scheduler.xml"
	filePath := strings.Split(wgetUrl, "/")
	ret_file := filePath[len(filePath)-1]
	fmt.Println(strings.Join(filePath[:len(filePath)-1], "/"))
	fmt.Println(ret_file)

	// 第一次增加
	//fmt.Println(ret_file)
	// 第二次增加
	//fmt.Println(ret_file)
	// 第三次增加
	//fmt.Println(ret_file)

	// ---
	//s := " 100 Mb/s"
	s := "Speed: Unknown!"

	fields := strings.Split(strings.TrimSpace(s), ":")
	if len(fields) != 2 {
		fmt.Println("len(fields)!=2")
		os.Exit(-1)
	}
	if strings.ToLower(strings.TrimSpace(fields[0])) != "speed" {
		fmt.Println("TrimSpace(fields[0])!=speed")
		os.Exit(-2)
	}

	compile := regexp.MustCompile("^[[:space:]]*([[:digit:]]+)[[:space:]]*Mb/s[[:space:]]*$")
	fmt.Println("fields[1]=", fields[1])
	speed := compile.FindStringSubmatch(fields[1])
	fmt.Println("len(speed)= ", len(speed))
	fmt.Println("speed= ", speed)

	if len(speed) <= 1 {
		fmt.Println("len(speed)<= 1, will exist...")
		os.Exit(-3)
	}

}
