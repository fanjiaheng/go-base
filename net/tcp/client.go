package tcp

import (
	"fmt"
	"net"
	"time"
)

func MainClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:18999")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	defer conn.Close()
	for {
		_, err = conn.Write([]byte("hello world"))
		if err != nil {
			fmt.Println("send failed, err: ", err)
			return
		}

		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err: ", err)
			return
		}
		fmt.Println(string(buf[:n]))
		time.Sleep(time.Second * 3)
	}
}
