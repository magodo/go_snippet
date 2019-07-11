package cache

type cacheChain struct {
	caches []Cache
}

func ChainCache(caches ...Cache) Cache {
	return &cacheChain{caches}
}

func (c *cacheChain) Set(rec *Record) error {
	for _, c := range c.caches {
		if err := c.Set(rec); err != nil {
			return err
		}
	}
	return nil
}

func (c *cacheChain) Get(key string) (*Record, error) {
	missCaches := []Cache{}
	var rec *Record
	var err error

	for _, c := range c.caches {
		rec, err = c.Get(key)
		if err != nil {
			if err != ErrCacheMiss {
				return nil, err
			}
			missCaches = append(missCaches, c)
			continue
		}
		break
	}

	for _, c := range missCaches {
		if err := c.Set(rec); err != nil {
			return nil, err
		}
	}
	return rec, nil
}
