package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/encrypt"
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
	secret := "c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ=="
	an, err := base64.StdEncoding.DecodeString(secret)
	fmt.Println(string(an), err)

	key := "MuxiStudio203304"
	code := "for {go func(){time.Sleep(1*time.Hour)}()}"

	errcode, err := encrypt.AESEncryptOutInBase64([]byte(code), []byte(key))
	if err != nil {
		fmt.Println(err)
	}
	type Code struct {
		Content   string `json:"content"`
		ExtraInfo string `json:"extra_info"`
	}
	tmp := Code{Content: string(errcode)}
	change, err := json.Marshal(tmp)
	if err != nil {
		fmt.Println(err)
	}
	param := bytes.NewBuffer(change)
	url := "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate"
	req, err := http.NewRequest("PUT", url, param)
	fmt.Println(req)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("code", "yyj")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoieXlqIiwiaWF0IjoxNjM3MjAzODUyLCJuYmYiOjE2MzcyMDM4NTJ9.Jzty8lQi4jaQ8Kzuk22ztYLRKXyi6KYZytc8tRlO1zI")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	answer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(answer))
	defer resp.Body.Close()
}
