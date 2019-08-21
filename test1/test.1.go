package main

import (
	"fmt"
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
	fmt.Println(ret_file)
}
