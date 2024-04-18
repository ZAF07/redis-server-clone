package clients

// Data types
const (
	RESPTypeSimpleString = byte('+')
	RESPTypeBulkString   = byte('$')
	RESPTypeArray        = byte('*')
	RESPTypeInteger      = byte(':')
)

// Command types
const (
	PingCmd = "ping"
	EchoCmd = "echo"
)

// Data types concrete
type BulkString struct { // check the first element in the split byte arr to determine its type
	Value  []byte
	Length int
	Cap    int
}

var commands = map[string]Command{
	"ping": Command{Cmd: []byte("ping"), MinArgs: 0},
	"echo": Command{Cmd: []byte("echo"), MinArgs: 1},
}

// func to map a cmd to concrete command type
// func convertToCommand(c []byte) Command {

// }

type Request struct {
	Cmd    *Command
	Args   [][]byte
	Length int
}

type Command struct {
	Cmd     []byte
	MinArgs int
}

func (c Command) validate(args int) bool {
	// TODO: Also validate that the params are valid (check this in docs)
	return args >= c.MinArgs
}

type RedisArray struct {
	Length int
	Value  []byte
}
