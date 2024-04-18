package appservices

type RedisCore interface {
	Ping() []byte
	Echo(s []byte) []byte
}
