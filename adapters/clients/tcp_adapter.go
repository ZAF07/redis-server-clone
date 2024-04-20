package clients

import (
	"bytes"
	"fmt"

	core "github.com/codecrafters-io/redis-starter-go/core"
	"github.com/codecrafters-io/redis-starter-go/protocol"
	"github.com/codecrafters-io/redis-starter-go/protocol/parsers"
)

// TCPAdapter receives the request from clients, parses and validates the request into RESP protocol and calls the core services to execute the redis commands
type TCPAdapter struct {
	core core.RedisCoreServices
	// shoul also inject the parser. the parser has diff versions
	parser parsers.RespParser
}

/*
NewTCPAdapter returns a new instance of a client TCPAdapter
This adapter allows interaction with the associated core layer
*/
func NewTCPAdapter(c core.RedisCoreServices, p parsers.RespParser) *TCPAdapter {
	return &TCPAdapter{
		core:   c,
		parser: p,
	}
}

/*
TEST INPUT FOR THE ECHO STAGE 5:
$ redis-cli echo grape
Sent bytes: "*2\r\n$4\r\necho\r\n$5\r\ngrape\r\n"
Received bytes: "+grape\r\n"
Received RESP value: "grape

*/

/*
Adapt is the adapter method for a client adapter to translate the client request to a call to the core layer
It returns the results of the core later implemention
*/
func (t *TCPAdapter) Adapt(r []byte) ([]byte, error) {
	var res []byte
	req := t.parser.Parse(r)

	// validate the args
	if req.GetArgsLength() >= req.Cmd.MinArgs {
		// based on the cmd, call the core service
		switch {
		case bytes.EqualFold(req.Cmd.Cmd, []byte(protocol.CMDPING)):
			// Check length of given PING args
			argLen := req.GetArgsLength()
			minArgLen := req.Cmd.MinArgs

			// Ping without args
			if argLen == minArgLen {
				res = t.core.Ping(nil)
			}

			if argLen == 1 {
				res = t.core.Ping(req.Args[0].GetValue())
			}

			// Too many args
			if argLen > minArgLen {
				return protocol.ErrNumArgs, nil
			}

			// res := t.core.Ping(req.Args[0].GetValue())
			return res, nil

		case bytes.EqualFold(req.Cmd.Cmd, []byte(protocol.CMDECHO)):
			fmt.Printf("❓ calling core echo with -->%+v, %+v", req, string(req.Args[0].GetValue()))
			// ❓ calling core echo with -->{Cmd:0xc000058060 Args:[[103 114 97 112 101] []] Length:50}, grape
			args := []byte{}
			for _, val := range req.Args {
				args = append(args, val.GetValue()...)
			}
			res = t.core.Echo(args)
			return res, nil
		}
	}

	// TODO: Impplement proper error handling. Notice diff between Redis errors (NOT ENOUGH ARGS) and runtime or logic error
	return []byte(protocol.NULLBULKSTRINGV1), nil
}

/*
Some mock data to visualise
PING req: *1\r\n$4\r\nping\r\n
ECHO req: *2\r\n$4\r\necho\r\n$3\r\nhey\r\n
*/
// func (t *TCPAdapter) ParseResp(r []byte) dtos.Request {
// 	// TODO: Clean this implementation. naybe i need to strip the splitted byte just below because i am seeing an extra element at the end of the splitted byte arr
// 	reqData := bytes.Split(r, []byte("\r\n"))
// 	c := reqData[2]

// 	// TODO: This is incorrect. check how do i extract all arguments only
// 	a := reqData[len(reqData)-1]
// 	fmt.Println("❌ --> ", string(reqData[0]), reqData)
// 	//  ❌ -->  *2 [[42 50] [36 52] [101 99 104 111] [36 53] [103 114 97 112 101] []]

// 	// parse the req as per RESP protocol to extract cmd and args
// 	cmd, err := extractCmd(c)
// 	if err != nil {
// 		log.Printf("error in parsing request: %+v", err)
// 	}

// 	// validate the cmd and args
// 	cmd.Validate(len(a))

// 	if bytes.EqualFold(cmd.Cmd, []byte(constants.CMDECHO)) {
// 		return dtos.Request{
// 			Cmd:    cmd,
// 			Args:   reqData[len(reqData)-2:],
// 			Length: int(reqData[0][1]),
// 		}
// 	}

// 	return protocol.Request{
// 		Cmd:    cmd,
// 		Args:   reqData[3:],        // TODO: Check implementation
// 		Length: int(reqData[0][1]), // TODO: Need to convert to int
// 	}

// }

// func extractCmd(b []byte) (*protocol.Command, error) {
// 	if cmd, ok := protocol.Commands[string(b)]; ok {
// 		return &cmd, nil
// 	}
// 	return nil, errors.New("in valid command")
// }
