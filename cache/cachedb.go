package cache

import (
	"time"
)

//Structs
type Item struct {
	Object     interface{}
	Expiration int64
}
type Cache struct {
	*cache
}
type cache struct {
	defaultExpiration time.Duration
	items             map[string]Item
}

func (item Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	//If item has expired return true
	return time.Now().UnixNano() > item.Expiration
}

//Creates new cache
func New(defaultExpiration, cleanupInterval time.Duration) *Cache {
	items := make(map[string]Item)
	return newCache(defaultExpiration, items)
}
func newCache(de time.Duration, m map[string]Item) *Cache {
	if de == 0 {
		de = -1
	}
	c := &cache{
		defaultExpiration: de,
		items:             m,
	}
	Cch := Cache{c}

	return &Cch
}
