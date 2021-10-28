package cache

import (
	models "go-cache-mongo/model"
	"sync"
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
	mutex             sync.RWMutex
}

func (item Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	//If item has expired return true
	return time.Now().UnixNano() > item.Expiration
}

//Adds item to cache
//If exist item key, this item will update
func (c *cache) Set(data models.InMemData, i interface{}) error {
	key := data.Key
	c.mutex.Lock()
	c.items[key] = Item{
		Object: i,
	}
	c.mutex.Unlock()
	return nil
}

//Gets the item where has been to cache
func (c *cache) Get(data models.InMemData) (interface{}, bool) {
	key := data.Key
	c.mutex.RLock()
	// "Inlining" of get and Expired
	item, found := c.items[key]
	if !found {
		c.mutex.RUnlock()
		return nil, false
	}
	c.mutex.RUnlock()
	return item.Object, true
}

//Gets all items in cache
func (c *cache) Items() map[string]Item {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	m := make(map[string]Item, len(c.items))
	for c, t := range c.items {
		m[c] = t
	}
	return m
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
