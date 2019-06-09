package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	timeStr := time.Now().Format("20060102")
	fmt.Println(timeStr)
	url := "http://172.16.189.124:9999/hadoop/hadoop_2.7.1_1.0_2018120518-20181205190803-jdk1.7.tar.gz"
	arr := strings.Split(url, "/")
	fmt.Println(arr[len(arr)-1])

}
