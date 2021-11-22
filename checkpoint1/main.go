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

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("code", "yyj")
	if err != nil {
		fmt.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	defer resp.Body.Close()
}
