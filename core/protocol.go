package core

/*
Request represents the structure of the incoming request.
The core consumes the Request struct populated with data from
*/
type Request struct {
	Cmd    *Command
	Args   [][]byte
	Length int
}

type Command struct {
	Cmd     []byte
	MinArgs int
}

const (
	RedisArrayType = '*'
)

var (
	PingCmd = []byte("ping")
	EchoCmd = []byte("echo")
)

type RedisArray struct {
	Length int
	Value  []byte
}
