package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	fmt.Printf("Server listening on port 4221\n")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		// start a goroutine to handle the new connection
		go handleRequest(conn)
	}
}
func handleRequest(conn net.Conn) {
	data := make([]byte, 4096)
	_, err := conn.Read(data)
	if err != nil {
		fmt.Printf("Error reading: %s\n", err)
		os.Exit(1)
	}
	lines := strings.Split(string(data), "\r\n")
	firstLineParts := strings.Split(lines[0], " ")
	method := firstLineParts[0]
	path := firstLineParts[1]
	userAgent := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "User-Agent: ") {
			userAgent = strings.TrimPrefix(line, "User-Agent: ")
			userAgent = strings.TrimSpace(userAgent)
			break
		}
	}
	if method != "GET" {
		fmt.Println("Only GET is supported")
		os.Exit(1)
	}
	if path == "/" {
		okRes := createHttpResp(200, []byte(""), "text/plain")
		conn.Write([]byte(okRes))
	} else if strings.HasPrefix(path, "/echo/") {
		suffix := strings.TrimPrefix(path, "/echo/")
		okRes := createHttpResp(200, []byte(suffix), "text/plain")
		conn.Write([]byte(okRes))
	} else if path == "/user-agent" {
		okRes := createHttpResp(200, []byte(userAgent), "text/plain")
		conn.Write([]byte(okRes))
	} else {
		notFoundRes := createHttpResp(404, []byte(""), "text/plain")
		conn.Write([]byte(notFoundRes))
	}
	conn.Close()
}
func createHttpResp(status int, body []byte, contentType string) string {
	resp := "HTTP/1.1 "
	switch status {
	case 200:
		resp += fmt.Sprintf("200 OK\r\n")
	case 404:
		resp += fmt.Sprintf("404 Not Found\r\n")
	}
	if len(body) > 0 {
		resp += fmt.Sprintf("Content-Type: %s\r\n", contentType)
		resp += fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(body), string(body))
	} else {
		resp += fmt.Sprintf("Content-Length: 0\r\n\r\n")
	}
	return resp
}
