package core

type RedisCoreServices interface {
	Ping(arg []byte) []byte
	Echo(arg ...[]byte) []byte
}
