package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// // You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)

	}

	time.Sleep(100 * time.Millisecond)
	defer conn.Close()

	reader := bufio.NewReader(conn)

	requestLine, _, err := reader.ReadLine()
	if err != nil {
		fmt.Print("error reading request line: ", err.Error())
		return
	}
	parts := strings.Split(string(requestLine), " ")
	if len(parts) >= 3 {
		method := parts[0]      // "GET"
		url := parts[1]         // "/hello"
		httpVersion := parts[2] // "HTTP/1.1"

		fmt.Printf("Method: %s\n", method)
		fmt.Printf("URL: %s\n", url)
		fmt.Printf("HTTP Version: %s\n", httpVersion)

		// Check if we have a valid URL (not empty or just a space)
		if url != "" && url != " " {
			response := "HTTP/1.1 200 OK\r\n\r\n"
			_, err = conn.Write([]byte(response))
			if err != nil {
				fmt.Println("Error writing response:", err.Error())
				os.Exit(1)
			}
		} else {
			// Send error response for invalid URLs
			response2 := "HTTP/1.1 400 Bad Request\r\n\r\n"
			_, err = conn.Write([]byte(response2))
			if err != nil {
				fmt.Println("Error writing error response:", err.Error())
				os.Exit(1)
			}
		}
	}

}
