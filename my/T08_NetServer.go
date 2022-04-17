package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888") // 相当于Java里new出来了ServerSocket
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("A client just connected")
		go handle(conn) // 这里go这个关键字关键，并发启动新协程（go routine）太简单了
	}
}

// 客户端连上来之后，每秒钟写给他一个"hello"
func handle(conn net.Conn) {
	defer conn.Close() // 不论写在什么位置最后都给他关掉
	for {
		_, err := io.WriteString(conn, "hello")
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
	}
}
