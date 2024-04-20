package dtos

import (
	"github.com/codecrafters-io/redis-starter-go/protocol/zredis"
)

/*
This is the transfer object.
Ther parser is responsible of taking the raw req and serialising it into a DTO
The parser then returns the Request object to the adapter

The adapter then
*/

type Command struct {
	Cmd     []byte
	MinArgs int
}

type Request struct {
	Cmd  *Command
	Args []zredis.RedisDataType
	// Length int // do i need?
}

func (r Request) GetArgsLength() int {
	return len(r.Args)
}

// TODO: Move to appropriate package
// type BulkString struct {
// 	Len   int
// 	Cap   int
// 	Value bytes.Buffer
// }

// func (b *BulkString) GetLength() int {
// 	return b.Value.Len()
// }

// func (b *BulkString) GetValue() []byte {
// 	return b.Value.Bytes()
// }
