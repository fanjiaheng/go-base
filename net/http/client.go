package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func MainClient() {
	resp, err := http.Get("http://127.0.0.1:9090")
	if err != nil {
		fmt.Println("get failed, err: ", err)
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
