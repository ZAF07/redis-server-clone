package clients

import (
	"bytes"
	"errors"
	"log"

	services "github.com/codecrafters-io/redis-starter-go/core/app_services"
)

// TCPAdapter receives the request from clients, parses and validates the request into RESP protocol and calls the core services to execute the redis commands
type TCPAdapter struct {
	core services.RedisCore
}

/*
NewTCPAdapter returns a new instance of a client TCPAdapter
This adapter allows interaction with the associated core layer
*/
func NewTCPAdapter(c services.RedisCore) *TCPAdapter {
	return &TCPAdapter{
		core: c,
	}
}

/*
Adapt is the adapter method for a client adapter to translate the client request to a call to the core layer
It returns the results of the core later implemention
*/
func (t *TCPAdapter) Adapt(r []byte) ([]byte, error) {

	// extract the cmd and args
	req := t.ParseResp(r)
	// based on the cmd, call the core service
	switch {
	case bytes.EqualFold(req.Cmd.Cmd, []byte(PingCmd)):
		res := t.core.Ping()
		return res, nil

	case bytes.EqualFold(req.Cmd.Cmd, []byte(EchoCmd)):
		res := t.core.Echo(req.Args[0])
		return res, nil
	}
	return []byte{}, nil
}

/*
PING req: *1\r\n$4\r\nping\r\n
ECHO req: *2\r\n$4\r\necho\r\n$3\r\nhey\r\n
*/
func (t *TCPAdapter) ParseResp(r []byte) Request {
	reqData := bytes.Split(r, []byte("\r\n"))
	c := reqData[2]
	// TODO: This is incorrect. check how do i extract all arguments only
	a := reqData[4]

	// parse the req as per RESP protocol to extract cmd and args
	cmd, err := extractCmd(c)
	if err != nil {
		log.Printf("error in parsing request: %+v", err)
	}

	// validate the cmd and args
	cmd.validate(len(a))

	return Request{
		Cmd:    cmd,
		Args:   reqData[3:],
		Length: int(reqData[0][1]),
	}

}

func extractCmd(b []byte) (*Command, error) {
	if cmd, ok := commands[string(b)]; ok {
		return &cmd, nil
	}
	return nil, errors.New("in valid command")
}
