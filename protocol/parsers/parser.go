package parsers

import (
	"github.com/ZAF07/redis-server-clone/dtos"
)

// Expose interface
type RespParser interface {
	Parse(r []byte) dtos.Request
}

// dont need here (should go to Redis layer)
var Commands = map[string]dtos.Command{
	"ping": {Cmd: []byte("ping"), MinArgs: 0},
	"echo": {Cmd: []byte("echo"), MinArgs: 1},
	"get":  {Cmd: []byte("get"), MinArgs: 1},
	"set":  {Cmd: []byte("set"), MinArgs: 2},
}
