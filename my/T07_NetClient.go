package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8888")
	fmt.Println(conn)
	log.Println("connected...")
	defer conn.Close()
	io.Copy(os.Stdout, conn)
}
