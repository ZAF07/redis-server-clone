package protocol

import "bytes"

// Data types
const (
	// RESP DATA TYPE
	RESPTypeSimpleString = byte('+')
	RESPTypeBulkString   = byte('$')
	RESPTypeArray        = byte('*')
	RESPTypeInteger      = byte(':')

	// // RESP response constants
	// PINGRESPV1       = "+PONG\r\n"
	// OKRESPV1         = "+OK\r\n"
	// NULLBULKSTRINGV1 = "$-1\r\n"

	// REDIS command constants
	CMDECHO = "echo"
	CMDPING = "ping"
	CMDSET  = "set"
	CMDGET  = "get"
)

var (

	// RESP response constants
	PINGRESPV1       = []byte("+PONG\r\n")
	OKRESPV1         = []byte("+OK\r\n")
	NULLBULKSTRINGV1 = []byte("$-1\r\n")

	// TODO: Implement better redis error. For example, an Unknown commad error also returns the given command
	ErrNumArgs    = []byte("(error) ERR wrong number of arguments for command")
	ErrUnknownCmd = []byte("(error) unknown command")
)

// interface representing a redis custom data type
type RedisDataType interface {
	GetLength() int
	GetValue() []byte
}

type BulkString struct { // check the first element in the split byte arr to determine its type
	Value  bytes.Buffer
	Length int
	Cap    int
}

func (b *BulkString) GetLength() int {
	return b.Value.Len()
}

func (b *BulkString) GetValue() []byte {
	return b.Value.Bytes()
}
