package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func Error(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	/*req, err := httptool.NewRequest(
		"",
		"",
		"",
		httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(req)

	// write your code below
	// ...*/
	url := "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination"
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	contenttype := writer.FormDataContentType()
	formfile, err := writer.CreateFormFile("file", "permute.go")
	Error(err)
	file, err := os.Open("./test.go")
	Error(err)
	defer file.Close()
	_, err = io.Copy(formfile, file)
	Error(err)
	writer.Close()
	req, err := http.NewRequest("POST", url, buf)
	Error(err)
	req.Header.Add("code", "yyj")
	req.Header.Add("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoieXlqIiwiaWF0IjoxNjM3MjAzODUyLCJuYmYiOjE2MzcyMDM4NTJ9.Jzty8lQi4jaQ8Kzuk22ztYLRKXyi6KYZytc8tRlO1zI")
	req.Header.Set("Content-Type", contenttype)

	resp, err := http.DefaultClient.Do(req)
	Error(err)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

}
