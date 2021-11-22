package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	client := &http.Client{}
	url := "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/code"

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("code", "yyj")

	req.Header.Add("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoieXlqIiwiaWF0IjoxNjM3MjA0MDMzLCJuYmYiOjE2MzcyMDQwMzN9.CTGHa5Pq9yAICFsVTC1Xq9dAl3k221j2ERZ2crGlZQ4")
	if err != nil {
		fmt.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content), err)
	defer resp.Body.Close()
}
