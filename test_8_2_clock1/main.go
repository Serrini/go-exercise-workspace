package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)
func main() {

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		//handleConn(conn)	// go handleConn(conn) 多个客户端就可以同时接收到时间了
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("Mon Jan 2 15:04:05\n")) // 格式化日期和时间
		io.WriteString(os.Stdout, "hhhhhhhhh")

		if err != nil {
			log.Print()
			return
		}

		time.Sleep(1 * time.Second)
	}
}
