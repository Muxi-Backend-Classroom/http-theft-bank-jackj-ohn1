package main

import (
	"fmt"
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
	if err != nil {
		fmt.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	passport := resp.Header.Get("passport")
	fmt.Println(passport)
	defer resp.Body.Close()
}
