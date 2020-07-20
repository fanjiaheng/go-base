package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func MainNewRequestClient() {

	req, err := http.NewRequest("get", "http://127.0.0.1:9090", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read from resp.Body failed, err: ", err)
		return
	}

	fmt.Println(string(body))
}
