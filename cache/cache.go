package cache

import (
	"errors"
	"time"
)

var (
	ErrCacheMiss = errors.New("cache miss")
)

type Cache interface {
	Set(*Record) error
	Get(key string) (*Record, error)
}

type Record struct {
	Key    string
	Value  []byte
	Expiry time.Duration
}
