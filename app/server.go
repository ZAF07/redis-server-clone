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
)

func main() {
	// Test
	inMemoryStore := core.NewInMemoryStore()
	core := core.NewRedisCore(inMemoryStore)
	parser := parsers.NewRESPParserV1()
	tcpAdapter := clients.NewTCPAdapter(core, parser)
	tcpServer := network.NewTCPServer("0.0.0.0", tcpAdapter, network.WithReadDeadline("25s"), network.WithReadDeadline("25s"))

	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("🚨 --> Logs from your program will appear here!")
	tcpServer.Start()
}
