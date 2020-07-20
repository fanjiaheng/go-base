package http

import (
	"fmt"
	"net/http"
)

func MainHttpServer() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
