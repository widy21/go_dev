package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}
	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	return nil
}

// sample usage
func main() {
	if len(os.Args) == 1 {
		fmt.Println("there is no arg...")
		return
	}
	var inputfilename string
	flag.StringVar(&inputfilename, "f", "", "no balncer name after [-f]")
	flag.Parse()
	fmt.Printf("[-b] inputfilename=%s \n", inputfilename)
	target_url := "http://commoncfg.jd.com/upload"
	//filename := "/Users/huaxiao/Documents/workspace/go/src/go_dev/Alexroom.jpg"
	//filename := "/Users/huaxiao/Documents/workspace/go/src/go_dev/json-lib-2.4-jdk15.jar"
	postFile(inputfilename, target_url)
}
