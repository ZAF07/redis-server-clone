package core

import "github.com/codecrafters-io/redis-starter-go/protocol/zredis"

type RedisCoreServices interface {
	Ping(arg []byte) []byte
	Echo(arg ...[]byte) []byte
	Set(k []byte, v zredis.RedisDataType) []byte
	Get(k []byte) []byte
}
