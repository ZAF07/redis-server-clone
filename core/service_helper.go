package core

import (
	"bytes"
	"strconv"
)

/*
formatResponse takes in an array of []byte and formats it into a RESP response
It appends '+' to the start and '\r\n' to the end
*/
func formatResponse(b ...[]byte) []byte {
	buf := bytes.NewBuffer(nil)
	buf.WriteByte('+')

	if len(b) == 1 {
		buf.Write(b[0])
		buf.WriteString("\r\n")
		return buf.Bytes()
	}

	for _, val := range b {
		buf.Write(val)
		// or
		// for _, v := range val {
		// 	buf.WriteByte(v)
		// }
	}

	buf.WriteString("\r\n")
	return buf.Bytes()
}

func responseBulkString(r []byte) []byte {
	// $3\r\nbar\r\n
	// Convert to decimal string representation
	bulkStrLen := strconv.AppendInt([]byte{}, int64(len(r)), 10) // Base 10 for decimal
	// Convert the string to a byte slice
	buf := bytes.NewBuffer(nil)
	buf.WriteByte('$')
	buf.Write([]byte(bulkStrLen))
	buf.WriteString("\r\n")
	buf.Write(r)
	buf.WriteString("\r\n")
	return buf.Bytes()
}
