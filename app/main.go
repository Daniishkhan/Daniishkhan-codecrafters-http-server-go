package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	// defer conn.Close() // Always clean up when done

	// Your existing parsing logic goes here
	// Read request, parse headers, send response

	fmt.Printf("🏁 Connection handled\n")
}
apackage main
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
func main() {
	listener, err := net.Listen("tcp", ":4221")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Printf("🚀 Server listening on :4221\n")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("❌ Accept error: %v\n", err)
			continue
		}

		fmt.Printf("📞 New connection accepted\n")

		// The magic line - handle each connection concurrently
		go handleConnection(conn)
	}
}
