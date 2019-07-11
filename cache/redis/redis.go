package redis

import (
	"sync"
	"time"

	"cache"

	"github.com/gomodule/redigo/redis"
)

type redisPool struct {
	sync.Mutex
	pool map[string]*redisCache
}

type redisCache struct {
	redis.Conn
}

var redispool = redisPool{
	pool: map[string]*redisCache{},
}

func NewCache(addr, password string) cache.Cache {
	// use addr as key in pool
	redispool.Lock()
	defer redispool.Unlock()
	if c, ok := redispool.pool[addr]; ok {
		return c
	}

	conn, err := redis.Dial("tcp", addr, redis.DialPassword(password))
	if err != nil {
		return nil
	}
	c := &redisCache{conn}
	redispool.pool[addr] = c
	return c
}

func (c *redisCache) Set(rec *cache.Record) error {
	if _, err := c.Do("SET", rec.Key, string(rec.Value)); err != nil {
		return err
	}
	if rec.Expiry > time.Duration(0) {
		if _, err := c.Do("EXPIRE", rec.Key, rec.Expiry.Seconds()); err != nil {
			return err
		}
	}
	return nil
}

func (c *redisCache) Get(key string) (*cache.Record, error) {
	v, err := redis.String(c.Do("GET", key))
	if err != nil {
		if err == redis.ErrNil {
			return nil, cache.ErrCacheMiss
		}
		return nil, err
	}
	ttl, err := redis.Int(c.Do("TTL", key))
	if err != nil {
		if err == redis.ErrNil {
			return nil, cache.ErrCacheMiss
		}
		return nil, err
	}
	return &cache.Record{
		Key:    key,
		Value:  []byte(v),
		Expiry: time.Duration(ttl) * time.Second,
	}, nil
}
