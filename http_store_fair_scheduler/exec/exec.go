package exec

import (
	"bufio"
	"log"
	"os"
	"strings"
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

func StoreFile(filePath, xmlStr string) int {

	// 如果文件夹不存在，则返回错误
	arr := strings.Split(filePath, "/")
	fileName := filePath[len(arr)-1]
	dir := strings.Join(arr[:len(arr)-1], "/")
	log.Println("fileName = ", fileName)
	log.Println("dir = ", dir)
	if !Exists(dir) {
		log.Println("dir is not existt ")
		return defaultFailedCode
	}

	file, error := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 755)
	if error != nil {
		log.Println("op file error: ", error)
	}

	writer := bufio.NewWriter(file)
	writer.WriteString(xmlStr)
	writer.Flush()

	return 0
}
