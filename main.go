package main

import (
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Processing the request")
	time.Sleep(5 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello Welcome to Backend World!\r\n"))
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}
	for {
		log.Println("Waiting a client to connect!")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Client Connected")
		go do(conn)
	}
}
