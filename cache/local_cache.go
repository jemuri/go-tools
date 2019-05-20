package cache

import (
	"sync"
	"time"
)

func init() {
	go autoRemoveExpiredKey()
}

var globalMap = make(map[string]Item)
var doubleCheck sync.RWMutex

var singleton = &localCache{
	items: globalMap,
}

//Cache ...
type Cache interface {
	Get(key string) interface{}
	Set(key string, value interface{}, duration ...time.Duration)
	SetEx(key string, value interface{}, duration time.Duration)
	Del(key string)
}

type Item struct {
	Data        interface{}
	ExpiredTime int64 //Nano
}

type localCache struct {
	items map[string]Item
	mutex sync.RWMutex
}

//NewLocalCache new a instance
func NewLocalCache() Cache {
	if singleton == nil {
		doubleCheck.Lock()
		defer doubleCheck.Unlock()
		if singleton == nil {
			singleton = &localCache{
				items: globalMap,
			}
		}
	}

	return singleton
}

//Get ...
func (c *localCache) Get(key string) interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	value, ok := c.items[key]
	if ok {
		if !value.isExpired() {
			return value.Data
		}
		delete(c.items, key)
	}
	return nil
}

//Set ...
func (c *localCache) Set(key string, value interface{}, duration ...time.Duration) {
	if len(duration) > 0 && duration[0] > 0 {
		c.SetEx(key, value, duration[0])
		return
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.items[key] = Item{
		Data:        value,
		ExpiredTime: -1,
	}
}

//setEx ...
func (c *localCache) SetEx(key string, value interface{}, duration time.Duration) {

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.items[key] = Item{
		Data:        value,
		ExpiredTime: time.Now().Add(duration).UnixNano(),
	}
}

//Del ...
func (c *localCache) Del(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.items, key)
}

//autoRemoveExpiredKey ...
func autoRemoveExpiredKey() {

}

//isExpired ...
func (i Item) isExpired() bool {
	if i.ExpiredTime == -1 {
		return false
	}
	return time.Now().UnixNano() > i.ExpiredTime
}
