package md5file

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Md5File(file_path string, file_name string) {
	// md5加密源文件
	abs_file_path := fmt.Sprintf("%s/%s", file_path, file_name)
	md5_file_path := fmt.Sprintf("%s/%s.md5", file_path, file_name)
	fmt.Println("abs_file_path = ", abs_file_path)
	fmt.Println("md5_file_path = ", md5_file_path)
	f, err := os.Open(abs_file_path)
	if err != nil {
		fmt.Printf("Error :%s\n", err.Error())
		panic(fmt.Sprintf("Error :%s\n", err.Error()))
		// return
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		fmt.Printf("Error :%s\n", err.Error())
		panic(fmt.Sprintf("Error :%s\n", err.Error()))
		// return
	}
	md5_str := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(md5_str)

	// 结果写入目标文件
	f1, err := os.Create(md5_file_path)
	check(err)
	defer f1.Close()
	_, err1 := f1.WriteString(md5_str)
	check(err1)
	f1.Sync()
}

func main() {
	Md5File("/tmp/testweb", "md5sha1sum-0.9.5.tar.gz")
}
