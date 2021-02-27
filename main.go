package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	if err := serve(); err != nil {
		panic("ERROR")
	}
}

func serve() error {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	conn, err := listen.Accept()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	s := bufio.NewScanner(conn)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		fmt.Println(s.Text())
	}

	if s.Err() != nil {
		fmt.Println(s.Err())
	}

	_, err = io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	if err != nil {
		log.Println("error")
	}

	return nil
}
