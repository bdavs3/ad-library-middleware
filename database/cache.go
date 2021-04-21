package database

import (
	"errors"
	"fmt"

	"github.com/patrickmn/go-cache"
)

// A CacheLayer may store a number of caches, each representing a certain database table.
type CacheLayer struct {
	Tables map[string]*cache.Cache
}

// NewCacheLayer returns a new CacheLayer instance without any tables.
func NewCacheLayer() *CacheLayer {
	return &CacheLayer{Tables: nil}
}

// AddTable adds a table with the given name to the CacheLayer.
func (cl *CacheLayer) AddTable(table string) {
	c := cache.New(cache.NoExpiration, cache.NoExpiration)
	cl.Tables[table] = c
}

// AddKVPair adds the given key/value pair to the given table in the CacheLayer.
func (cl *CacheLayer) AddKVPair(table, key string, value interface{}) error {
	t, ok := cl.Tables[table]
	if !ok {
		return errors.New(fmt.Sprintf("table %v not found in the cache layer", table))
	}

	t.Set(key, value, cache.NoExpiration)

	return nil
}

// FindValue looks for the given value in the given table in the CacheLayer.
func (cl *CacheLayer) FindValue(table, key string) (interface{}, error) {
	t, ok := cl.Tables[table]
	if !ok {
		return nil, errors.New(fmt.Sprintf("table %v not found in the cache layer", table))
	}

	value, found := t.Get(key)
	if found {
		return value, nil
	}

	return nil, nil
}
