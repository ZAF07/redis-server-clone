package network

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

type TCPServerOptions func(t *TCPServer)

type TCPServer struct {
	host          string
	port          string
	readDeadline  time.Duration
	writeDeadline time.Duration
	adapter       TCPClientAdapter
}

/*
NewTCPServer takes in server config options and returns a new instance of a TCP server.
*/
func NewTCPServer(host string, c TCPClientAdapter, opts ...TCPServerOptions) *TCPServer {
	tcpServer := &TCPServer{
		host:    host,
		port:    "6379",
		adapter: c,
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			opt(tcpServer)
		}
	}

	return tcpServer
}

func (t *TCPServer) getTCPAddress() string {
	return fmt.Sprintf("%s:%s", t.host, t.port)
}

func (t *TCPServer) Start() {
	l, err := net.Listen("tcp", t.getTCPAddress())
	if err != nil {
		fmt.Printf("Failed to bind to port '%s'", t.port)
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		// For each connection, spin up a handle() routine representing a single unique connection
		go t.handle(conn)
	}
}

func (t *TCPServer) handle(conn net.Conn) {
	defer conn.Close()

	// Check why i always see 1024 as byte array size
	buf := make([]byte, 1024)
	for {
		// Timeout for read
		/*
			SetReadDeadline is typically set for each new Read. Not the connection itself.
			So lets say you set the timeout to 5Sec but the entire read would take 6Sec, the current read process would fail with a timeout
			But the connection would still stay open for subsequent client requests
		*/
		tErr := conn.SetReadDeadline(time.Now().Add(t.readDeadline))
		if tErr != nil {
			fmt.Println("🚨 Read timeout --> ", tErr)
		}

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("🚨 error reading from client: %v\n", err)
		}
		fmt.Printf("💡 Message from client: %v\n", string(buf[:n]))

		// Parse the resquest to extract the command and arguments to pass to the client adapter
		// cmd, args, err := t.parseRequest(buf[:n])
		// if err != nil {
		// 	log.Fatalf("error parsing request: %+v", err)
		// }

		// Call the client adapter passing the command and arguments. Returns a response from the core logic which is the result of the request
		var res []byte
		if len(buf[:n]) > 0 {

			res, err = t.adapter.Adapt(buf[:n])
			if err != nil {
				log.Fatalf("error adapting to core layer: %+v", err)
			}
		}

		// conn.Write([]byte("+PONG\r\n"))
		conn.Write(res)
	}
}

// TODO: Parse should be a part of the core service. TCP server accepts requests, calls TCP adapter which then calls the core service to parse the request. it gets back the cmd and args, which TCP adapter then adapts and calls the core services
func (t *TCPServer) parseRequest(r []byte) (string, []string, error) {
	// TODO: Implement the parsing of the RESP string to get the cmd and args (if any)
	return "command", []string{}, nil
}

func WithReadDeadline(rd string) TCPServerOptions {
	rdl, err := time.ParseDuration(rd)
	if err != nil {
		log.Fatalf("error parsing read deadline while initialising TCP server: %+v", err)
	}
	return func(t *TCPServer) {
		t.readDeadline = rdl
	}
}

func WithWriteDeadline(wd string) TCPServerOptions {
	wdl, err := time.ParseDuration(wd)
	if err != nil {
		log.Fatalf("error parsing write deadline while initialising TCP server: %+v", err)
	}
	return func(t *TCPServer) {
		t.writeDeadline = wdl
	}
}

func WithPort(p string) TCPServerOptions {
	port := p
	if _, err := strconv.Atoi(p); err != nil {
		log.Printf("given port address is invalid. default port 6379 is being used. given: %s. error: %+v", p, err)
		port = "6379"
	}
	return func(t *TCPServer) {
		t.port = port
	}
}
