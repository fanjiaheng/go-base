package tcp

import (
	"bufio"
	"fmt"
	"net"
)

type TcpServer struct {
	Ip   string
	Port string
}

// Tcp服务器主程序
func (this *TcpServer) MainTcpServer() {

	fmt.Println(this.Ip + ":" + this.Port)

	listen, err := net.Listen("tcp", this.Ip+":"+this.Port)
	if err != nil {
		fmt.Println("listen failed.")
		return
	}

	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err: ", err)
			continue
		}

		addr := conn.RemoteAddr()
		fmt.Println("接收到客户端的信息: ", addr)
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close() // 关闭连接

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read client failed, err: ", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到客户端发送过来的数据: ", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据到客户端
	}
}
