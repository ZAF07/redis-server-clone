package core

import (
	"bytes"
	"log"
)

// Need to implement the parser now
/*
RedisParser takes in the RESP serialised request sent from the client
It parses the RESP string and validates and extracts the command and arguments.
It returns a Request object
*/

type RedisParser struct{}

func NewRedisParser() *RedisParser {
	return &RedisParser{}
}

/*
PING req: *1\r\n$4\r\nping\r\n
ECHO req: *2\r\n$4\r\necho\r\n$3\r\nhey\r\n
*/
func (p RedisParser) ParseReq(r []byte) *Request {
	// 💡 NOTE: For now, assume that all request from clients comes in RESP array form
	// Some commands from the client are sent using a simple string
	// Majority of the commands are sent using a RESP array

	// 💡DO LATER: Step 1: Validate that the request is in correct format

	pReq := bytes.Split(r, []byte("\r\n"))

	cmd, err := extractCmd(pReq[2])
	if err != nil {
		log.Fatalf("error extracting command in core parser. error: %+v", err)
	}

	req := &Request{
		Cmd:    &cmd,
		Args:   pReq[1:],
		Length: int(pReq[0][1]),
	}

	return req
}

func extractCmd(r []byte) (Command, error) {
	//TODO: to check if the given request is a RESP array or a simple string
	var cmd Command
	cmd.Cmd = r

	// Ping Cmd
	if bytes.EqualFold(r, PingCmd) {
		cmd.MinArgs = 0
	}

	if bytes.EqualFold(r, EchoCmd) {
		cmd.MinArgs = 1
	}

	return cmd, nil
}

func isCaseInsensitiveEqual(data1 []byte, data2 []byte) bool {
	if len(data1) != len(data2) {
		return false
	}
	for i := range data1 {
		diff := data1[i] ^ data2[i]
		// Check if the difference is only due to case (bits 0-5)
		if diff&(1<<5|1<<4|1<<3|1<<2|1<<1|1) != 0 {
			return false
		}
	}
	return true
}
