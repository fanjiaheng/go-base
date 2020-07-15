package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func MainClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:18999")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n') //读入用户输入
		inputInfo := strings.Trim(input, "\r\n")
		fmt.Println(inputInfo)
		if strings.ToUpper(inputInfo) == "Q" { //输入q退出
			return
		}

		_, err = conn.Write([]byte(inputInfo))
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
	}
}
