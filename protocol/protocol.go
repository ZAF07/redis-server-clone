package protocol

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
