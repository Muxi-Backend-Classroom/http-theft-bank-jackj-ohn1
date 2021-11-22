package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

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
	url := "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate"
	req, err := http.NewRequest("GET", url, nil)
	Error(err)
	req.Header.Add("code", "yyj")
	req.Header.Add("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoieXlqIiwiaWF0IjoxNjM3MjAzODUyLCJuYmYiOjE2MzcyMDM4NTJ9.Jzty8lQi4jaQ8Kzuk22ztYLRKXyi6KYZytc8tRlO1zI")
	resp, err := http.DefaultClient.Do(req)
	Error(err)
	answer, err := ioutil.ReadAll(resp.Body)
	Error(err)
	//fmt.Println(string(answer))
	defer resp.Body.Close()

	//下载图片
	url = "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/iris_sample"
	req, err = http.NewRequest("GET", url, nil)
	Error(err)

	//设置头
	req.Header.Add("code", "yyj")
	req.Header.Add("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoieXlqIiwiaWF0IjoxNjM3MjAzODUyLCJuYmYiOjE2MzcyMDM4NTJ9.Jzty8lQi4jaQ8Kzuk22ztYLRKXyi6KYZytc8tRlO1zI")

	//获取响应体
	resp, err = http.DefaultClient.Do(req)
	Error(err)
	_, err = ioutil.ReadAll(resp.Body)
	Error(err)

	//将body里的数据写入文件当中，写入之后要删除一些无用的东西，所以这里完成后后续就不要再进行了

	//读取文件中的内容，并进行解码
	tmp, err := ioutil.ReadFile("./虹膜.txt")
	//fmt.Println(len(tmp))
	Error(err)
	content, err := base64.StdEncoding.DecodeString(string(tmp))
	Error(err)

	//解码之后再写入
	_ = ioutil.WriteFile("./虹膜.jpg", content, 0777)

	defer resp.Body.Close()

	//上传图片
	url = "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate"

	buf := new(bytes.Buffer)                    //实例化一个结构体
	writer := multipart.NewWriter(buf)          //返回一个Writer指针
	ContentType := writer.FormDataContentType() //获取表单的提交格式
	form, err := writer.CreateFormFile("file", "虹膜.jpg")
	Error(err)

	file, err := os.Open("./虹膜.jpg")
	Error(err)
	defer file.Close()
	_, err = io.Copy(form, file)
	Error(err)
	writer.Close() //发送之前必须调用close（）以写入结尾行

	req, err = http.NewRequest("POST", url, buf)
	Error(err)

	req.Header.Set("Content-Type", ContentType)
	req.Header.Add("code", "yyj")
	req.Header.Add("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoieXlqIiwiaWF0IjoxNjM3MjAzODUyLCJuYmYiOjE2MzcyMDM4NTJ9.Jzty8lQi4jaQ8Kzuk22ztYLRKXyi6KYZytc8tRlO1zI")

	resp, err = http.DefaultClient.Do(req)
	Error(err)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
