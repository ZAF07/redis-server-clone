package network

/*
Commands comes into the server as a RESP array *<length>\r\n$<length>\r\n<command-in-character>\r\n<additional-types>
*/
type TCPClientAdapter interface {
	Adapt(r []byte) ([]byte, error)
}
