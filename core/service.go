package core

import (
	"bytes"
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/protocol"
)

// The in-memory storage implementation goes here as well
type RedisCore struct{}

func NewRedisCore() *RedisCore {
	return &RedisCore{}
}

/*
Ping command can have one optional arg
If no arg is given, it simple replies with 'PONG'
*/
func (r *RedisCore) Ping(arg []byte) []byte {
	fmt.Println("GOT IN CORE PING -> ")
	if arg != nil {
		return []byte(arg)
	}
	return protocol.PINGRESPV1
}

func (r *RedisCore) Echo(arg ...[]byte) []byte {
	fmt.Println("GOT IN CORE ECHO -> ", string(arg[0]))
	return formatResponse(arg...)
}

func formatResponse(b ...[]byte) []byte {
	buf := bytes.NewBuffer(nil)
	buf.WriteByte('+')

	if len(b) == 1 {
		buf.Write(b[0])
		buf.WriteString("\r\n")
		return buf.Bytes()
	}

	for _, val := range b {
		buf.Write(val)
		// or
		// for _, v := range val {
		// 	buf.WriteByte(v)
		// }
	}

	buf.WriteString("\r\n")
	return buf.Bytes()
}
