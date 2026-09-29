package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Listen failed: %v", err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Acceptance error: %v", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := bufio.NewReader(conn)

	for {
		line, err := buf.ReadString('\n')
		if err != nil || line == "\r\n" || line == "\n" {
			break
		}
	}

	body := "Hello you connected to a server"

	currentDate := time.Now().UTC().UTC().Format(http.TimeFormat)

	respose := fmt.Sprintf(
		"HTTP/1.0 200 OK\r\n"+
			"Date: %s\r\n"+
			"Content-Length: %d\r\n"+
			"Content-Type: text/plain; charset=utf-8\r\n"+
			"\r\n"+
			"%s",
		currentDate,
		len(body),
		body,
	)

	_, err := conn.Write([]byte(respose))
	if err != nil {
		log.Fatalf("write failed: %v", err)
	}
}
