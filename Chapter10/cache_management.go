///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing sync and time  package
import (
	"fmt"
	"sync"
	"time"
)

// CacheObject class
type CacheObject struct {
	Value      string
	TimeToLive int64
}

// IfExpired method
func (cacheObject CacheObject) IfExpired() bool {
	if cacheObject.TimeToLive == 0 {
		return false
	}
	return time.Now().UnixNano() > cacheObject.TimeToLive
}

//Cache class
type Cache struct {
	objects map[string]CacheObject
	mutex   *sync.RWMutex
}

//NewCache method
func NewCache() *Cache {
	return &Cache{
		objects: make(map[string]CacheObject),
		mutex:   &sync.RWMutex{},
	}
}

//GetObject method
func (cache Cache) GetObject(cacheKey string) string {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	var object CacheObject
	object = cache.objects[cacheKey]

	if object.IfExpired() {
		delete(cache.objects, cacheKey)
		return ""
	}
	return object.Value
}

//SetValue method
func (cache Cache) SetValue(cacheKey string, cacheValue string, timeToLive time.Duration) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	cache.objects[cacheKey] = CacheObject{
		Value:      cacheValue,
		TimeToLive: time.Now().Add(timeToLive).UnixNano(),
	}
}

// main method
func main() {

	var cache *Cache
	cache = NewCache()

	cache.SetValue("name", "john smith", 200000000)
	var name string
	name = cache.GetObject("name")
	fmt.Println(name)

}
