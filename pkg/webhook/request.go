package webhook

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func JsonPost(webhook_url string, request_body string) {
	//json序列化
	var jsonStr = []byte(request_body)

	req, err := http.NewRequest("POST", webhook_url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
