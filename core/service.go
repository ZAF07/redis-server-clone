package core

import "fmt"

type RedisCore struct{}

func NewRedisCore() *RedisCore {
	return &RedisCore{}
}

/*
Ping command can have one optional arg
If no arg is given, it simple replies with 'PONG'
*/
func (r *RedisCore) Ping() []byte {
	fmt.Println("GOT IN CORE PING -> ")
	// if arg != nil {
	// 	return arg
	// }
	return []byte("+PONG\r\n")
}

func (r *RedisCore) Echo(s []byte) []byte {
	fmt.Println("GOT IN CORE ECHO -> ", string(s))
	return s
}
