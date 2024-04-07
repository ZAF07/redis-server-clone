package main

import (
	"fmt"
	"time"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("🚨 --> Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)

	for {
		// Timeout for read
		tErr := conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		if tErr != nil {
			fmt.Println("🚨 Read timeout --> ", tErr)
		}
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("🚨 error reading from client: %v", err)
		}
		fmt.Printf("💡 Message from client: %v", buf[:n])

		conn.Write([]byte("+PONG\r\n"))
	}
	// write to conn
}
