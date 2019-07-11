package memcached

import (
	"sort"
	"strings"
	"sync"

	"cache"

	"github.com/bradfitz/gomemcache/memcache"
)

type memcachedPool struct {
	sync.Mutex
	pool map[string]*memcachedCache
}

type memcachedCache struct {
	*memcache.Client
}

var memcachedpool = memcachedPool{
	pool: map[string]*memcachedCache{},
}

func NewCache(servers ...string) cache.Cache {
	// construct the servers into a deterministic string
	sort.Strings(servers)
	serversStr := strings.Join(servers, "|")
	memcachedpool.Lock()
	defer memcachedpool.Unlock()
	if c, ok := memcachedpool.pool[serversStr]; ok {
		return c
	}
	c := &memcachedCache{memcache.New(servers...)}
	memcachedpool.pool[serversStr] = c
	return c
}

func (m *memcachedCache) Set(rec *cache.Record) error {
	item := memcache.Item{
		Key:        rec.Key,
		Value:      rec.Value,
		Flags:      0,
		Expiration: int32(rec.Expiry.Seconds()),
	}
	return m.Client.Set(&item)
}

func (m *memcachedCache) Get(key string) (*cache.Record, error) {
	item, err := m.Client.Get(key)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			return nil, cache.ErrCacheMiss
		}
		return nil, err
	}
	return &cache.Record{
		Key:   item.Key,
		Value: item.Value,
		// TODO: memcached not support getting expire time
		//Expiry: time.Second * time.Duration(item.Expiration),
	}, nil
}
