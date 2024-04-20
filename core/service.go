package core

import (
	"fmt"

	"github.com/ZAF07/redis-server-clone/protocol"
	"github.com/ZAF07/redis-server-clone/protocol/zredis"
)

// The in-memory storage implementation goes here as well
type RedisCore struct {
	storage InMemoryStore
}

func NewRedisCore(s InMemoryStore) *RedisCore {
	return &RedisCore{
		storage: s,
	}
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

func (r *RedisCore) Set(k []byte, v zredis.RedisDataType) []byte {
	fmt.Printf("❌service -> %s || %+v\n", string(k), v.GetValue())
	r.storage.Set(k, v)

	return protocol.OKRESPV1
}

func (r *RedisCore) Get(k []byte) []byte {
	res := r.storage.Get(k)
	if res != nil {
		return responseBulkString(res.GetValue())
	}

	return protocol.NULLBULKSTRINGV1
}
