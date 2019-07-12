package memory

import (
	"sync"
	"time"

	"cache"
)

type memCache struct {
	sync.RWMutex
	values map[string]*memRecord
}

func NewCache() cache.Cache {
	return &memCache{values: make(map[string]*memRecord)}
}

type memRecord struct {
	r *cache.Record
	c time.Time
}

func (s *memCache) Set(rec *cache.Record) error {
	s.Lock()
	defer s.Unlock()

	s.values[rec.Key] = &memRecord{
		r: rec,
		c: time.Now(),
	}
	return nil
}

func (s *memCache) Get(key string) (*cache.Record, error) {
	s.RLock()
	defer s.RUnlock()

	v, ok := s.values[key]
	if !ok {
		return nil, cache.ErrCacheMiss
	}

	d := time.Since(v.c)

	// check expiry
	if v.r.Expiry > time.Duration(0) {
		if d > v.r.Expiry {
			return nil, cache.ErrCacheMiss
		}
		// update expiry
		v.r.Expiry -= d
		v.c = time.Now()
	}

	return v.r, nil
}
