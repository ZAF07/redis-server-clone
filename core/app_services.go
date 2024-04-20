package core

import "github.com/ZAF07/redis-server-clone/protocol/zredis"

type RedisCoreServices interface {
	Ping(arg []byte) []byte
	Echo(arg ...[]byte) []byte
	Set(k []byte, v zredis.RedisDataType) []byte
	Get(k []byte) []byte
}
