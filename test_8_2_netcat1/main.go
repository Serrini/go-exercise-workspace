package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	go mustCOpy(os.Stdout, conn)
	mustCOpy(os.Stdout, conn) //将标准输⼊复制到server，客户端程序关闭标准输⼊时，后台goroutine可能依然在⼯作
}

func mustCOpy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
