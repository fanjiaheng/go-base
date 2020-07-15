package tcp

import (
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	fmt.Println("******Tcp服务器测试******")

	s := &TcpServer{
		Ip:   "127.0.0.1",
		Port: "18999",
	}

	s.MainTcpServer()
}
