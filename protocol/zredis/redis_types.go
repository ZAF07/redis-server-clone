package zredis

import "bytes"

// interface representing a redis custom data type
type RedisDataType interface {
	GetLength() int
	GetValue() []byte
}

// This should go to a separate package because the core also uses it
type BulkString struct { // check the first element in the split byte arr to determine its type
	Value  bytes.Buffer
	Length int
	Cap    int
}

func (b *BulkString) GetLength() int {
	return b.Value.Len()
}

func (b *BulkString) GetValue() []byte {
	return b.Value.Bytes()
}
