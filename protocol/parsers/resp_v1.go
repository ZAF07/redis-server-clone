package parsers

import (
	"bytes"
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/dtos"
	"github.com/codecrafters-io/redis-starter-go/protocol"
)

// TODO: Handle case for when ping is sent without arguments
// TODO: Implement error handling. There are two kinds if error. Redis specific errors with bytes response and logic or runtime error that needs to be resolved
// TODO: Implement Request struct pooling to reduce creation of a new Request struct each time

type RespParserV1 struct{}

func NewRESPParserV1() RespParserV1 {
	return RespParserV1{}
}

func (r RespParserV1) Parse(req []byte) dtos.Request {

	reqObj := dtos.Request{}

	// seperate the cmd and args based on the firstIdx
	splitCmdArgs := bytes.SplitAfter(req, []byte("\r\n"))
	fmt.Println("Sep --> ", splitCmdArgs)

	/*
		idx 0 = type & length of request
		idx 1 = type & length of command
		idx 2 = command
		idx 3: = arguments to the command
	*/
	c := splitCmdArgs[2]
	reqObj.Cmd = getCmd(c)

	a := splitCmdArgs[3 : len(splitCmdArgs)-1]
	fmt.Println("🚨 => ", a)

	if len(a) < 2 {
		return reqObj
	}
	reqObj.Args = getArgs(a)

	// reqObj := dtos.Request{
	// 	Cmd:  &cmd,
	// 	Args: args,
	// }

	return reqObj
}

func getCmd(r []byte) *dtos.Command {
	c := bytes.TrimRight(r, "\r\n")
	if cmd, ok := Commands[string(c)]; ok {
		return &cmd
	} else {
		return nil
	}
}

func getArgs(r [][]byte) []protocol.RedisDataType {
	args := make([]protocol.RedisDataType, len(r)/2)
	// fmt.Println("length --> ", len(args), string(r[0]))
	for i := 0; i < len(r); i += 2 {
		t := bytes.TrimRight(r[i], "\r\n")
		a := bytes.TrimRight(r[i+1], "\r\n")

		// figuriing out the type ($, +)
		val := bytes.NewBuffer(a)
		switch t[0] {
		case '$':
			dt := protocol.BulkString{
				Value:  *val,
				Length: int(t[1]),
				Cap:    val.Cap(),
			}
			args[i/2] = &dt
		}

	}
	return args
}
