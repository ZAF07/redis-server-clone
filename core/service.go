package core

type RedisCore struct{}

func NewRedisCore() *RedisCore {
	return &RedisCore{}
}

/*
Ping command can have one optional arg
If no arg is given, it simple replies with 'PONG'
*/
func (r *RedisCore) Ping(arg []byte) []byte {
	if arg != nil {
		return arg
	}
	return []byte("+PONG\r\n")
}

func (r *RedisCore) Echo(s []byte) []byte {
	return s
}
