package main

/*
TO SUBMIT:
git add .
git commit --allow-empty -m 'pass 2nd stage' # any msg
git push origin master
*/

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/adapters/clients"
	"github.com/codecrafters-io/redis-starter-go/adapters/network"
	"github.com/codecrafters-io/redis-starter-go/core"
	"github.com/codecrafters-io/redis-starter-go/protocol/parsers"
	// Uncomment this block to pass the first stage
)

func main() {
	core := core.NewRedisCore()
	parser := parsers.NewRESPParserV1()
	tcpAdapter := clients.NewTCPAdapter(core, parser)
	tcpServer := network.NewTCPServer("0.0.0.0", tcpAdapter, network.WithReadDeadline("25s"), network.WithReadDeadline("25s"))

	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("🚨 --> Logs from your program will appear here!")
	tcpServer.Start()

	// Uncomment this block to pass the first stage
	//
	// l, err := net.Listen("tcp", "0.0.0.0:6379")
	// if err != nil {
	// 	fmt.Println("Failed to bind to port 6379")
	// 	os.Exit(1)
	// }
	// defer l.Close()

	// for {
	// 	conn, err := l.Accept()
	// 	if err != nil {
	// 		fmt.Println("Error accepting connection: ", err.Error())
	// 		os.Exit(1)
	// 	}

	// 	go handler(conn)
	// }
}

// func handler(conn net.Conn) {
// 	defer conn.Close()
// 	buf := make([]byte, 1024)

// 	for {
// 		// Timeout for read
// 		/*
// 			SetReadDeadline is typically set for each new Read. Not the connection itself.
// 			So lets say you set the timeout to 5Sec but the entire read would take 6Sec, the current read process would fail with a timeout
// 			But the connection would still stay open for subsequent client requests
// 		*/
// 		tErr := conn.SetReadDeadline(time.Now().Add(20 * time.Second))
// 		if tErr != nil {
// 			fmt.Println("🚨 Read timeout --> ", tErr)
// 		}

// 		n, err := conn.Read(buf)
// 		if err != nil {
// 			fmt.Printf("🚨 error reading from client: %v\n", err)
// 		}
// 		fmt.Printf("💡 Message from client: %v\n", string(buf[:n]))

// 		conn.Write([]byte("+PONG\r\n"))
// 	}
// 	// write to conn
// }
