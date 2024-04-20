package core

import (
	"time"

	"github.com/ZAF07/redis-server-clone/protocol/zredis"
)

type InMemoryStore interface {
	Set(k []byte, v zredis.RedisDataType)
	Get(k []byte) zredis.RedisDataType
}

type InMemoryStoreV1 struct {
	items map[string]*values
	// expiration *ExpirationHeap
}

func NewInMemoryStore() InMemoryStore {
	return &InMemoryStoreV1{
		items: make(map[string]*values),
	}
}

type values struct {
	value   zredis.RedisDataType
	ttl     time.Duration
	expires time.Time
}

func (i *InMemoryStoreV1) Set(k []byte, v zredis.RedisDataType) {
	newValue := &values{
		value: v,
	}

	i.items[string(k)] = newValue
}

func (i *InMemoryStoreV1) Get(k []byte) zredis.RedisDataType {
	if kv, ok := i.items[string(k)]; ok {
		return kv.value
	}
	return nil
}
