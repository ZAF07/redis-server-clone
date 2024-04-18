package main

import (
	"database/sql"
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

		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)

	for {
		// Timeout for read
		/*
			SetReadDeadline is typically set for each new Read. Not the connection itself.
			So lets say you set the timeout to 5Sec but the entire read would take 6Sec, the current read process would fail with a timeout
			But the connection would still stay open for subsequent client requests
		*/
		tErr := conn.SetReadDeadline(time.Now().Add(20 * time.Second))
		if tErr != nil {
			fmt.Println("🚨 Read timeout --> ", tErr)
		}

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("🚨 error reading from client: %v\n", err)
		}
		fmt.Printf("💡 Message from client: %v\n", string(buf[:n]))

		conn.Write([]byte("+PONG\r\n"))
	}
	// write to conn
}

// Port for interactiom with the core layer
type CoreLogicPort interface {
	Get()
	Set()
	Del()
}

// Port for the core to communicate with external dependencies like a database
type Persistence interface {
	Save()
	Del()
}

type TCPServer struct {
	Adapter ClientAdapter
}

func (t *TCPServer) Start() {
	// Starts a tcp server and for each new connection, it parses the request and calls the adapter
	// Once it parses the request and gets the command that the client wants to execute, it calls the appropiate adapter methods

	// AdaptGet for GET request
	t.Adapter.AdaptGet()

	// AdaptSet for SET requests
	t.Adapter.AdaptSet()

}

// This could be a postgres, mysql or mongodb implementation. It is agnostic as long as it implements the persistence port
type PersistenceAdapter struct {
	db sql.Conn
}

func NewPersistenceAdapter() *PersistenceAdapter {
	return &PersistenceAdapter{
		db: sql.Conn{},
	}
}

// The actual implementation of the persistance layer
func (p *PersistenceAdapter) Save() {}
func (p *PersistenceAdapter) Del()  {}

// Adapter for clients to interact with the core
type ClientAdapter struct {
	core CoreLogicPort
}

func NewClientAdapter(c CoreLogicPort) *ClientAdapter {
	return &ClientAdapter{
		core: c,
	}
}

// Adapter's adapting the client request to the core logic
func (a *ClientAdapter) AdaptGet() {
	a.core.Get()
}
func (a *ClientAdapter) AdaptSet() {
	a.core.Set()
}

// The actual core layer. has methods that perform core business logic
// The core also has ports to interact with the outside like a database
type CoreLogic struct {
	Persistence Persistence
}

// Returns a new core layer (called upon app startup)
func NewCoreLogic(p Persistence) *CoreLogic {
	return &CoreLogic{
		Persistence: p,
	}
}

// Actual implementation of the core logic
func (c *CoreLogic) Get() {}
func (c *CoreLogic) Del() {}
func (c *CoreLogic) Set() {
	c.Persistence.Save()
}

func startApp() {
	externalDB := NewPersistenceAdapter()
	core := NewCoreLogic(externalDB)
	clientAdapter := NewClientAdapter(core)

	clientAdapter.AdaptGet()
}
