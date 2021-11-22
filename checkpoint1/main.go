package main

import (
	"fmt"
	"net/http"
        "io/ioutil"
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

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("code", "yyj")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
