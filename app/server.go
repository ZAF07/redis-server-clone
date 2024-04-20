package main

/*
TO SUBMIT:
git add .
git commit --allow-empty -m 'pass 2nd stage' # any msg
git push origin master
*/

import (
	"fmt"

	"github.com/ZAF07/redis-server-clone/adapters/clients"
	"github.com/ZAF07/redis-server-clone/adapters/network"
	"github.com/ZAF07/redis-server-clone/core"
	"github.com/ZAF07/redis-server-clone/protocol/parsers"
)

func main() {
	// TODO: Allow CLI args for users to input port and host
	inMemoryStore := core.NewInMemoryStore()
	core := core.NewRedisCore(inMemoryStore)
	parser := parsers.NewRESPParserV1()
	tcpAdapter := clients.NewTCPAdapter(core, parser)
	tcpServer := network.NewTCPServer("0.0.0.0", tcpAdapter, network.WithReadDeadline("25s"), network.WithReadDeadline("25s"))

	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("🚨 --> Logs from your program will appear here!")
	tcpServer.Start()
}
