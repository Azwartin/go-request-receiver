package server

import (
	"sort"
	"sync"
)

//storage is concurent safe storage for any struct
type storage struct {
	mutex sync.RWMutex
	items map[string]interface{}
}

var instance *storage
var once sync.Once

//GetStorage - get storage as singleton
func GetStorage() *storage {
	once.Do(func() {
		instance = &storage{
			mutex: sync.RWMutex{},
			items: map[string]interface{}{},
		}
	})

	return instance
}

//Load returns value stored in the map for a key, or nil if no value present
//The ok result indicates whether value was found in the map.
func (s *storage) Load(key string) (interface{}, bool) {
	value, ok := s.items[key]
	return value, ok
}

//Load returns slice with values stored in the map in order with offset and limit
//This is not a production ready code, just for simplifying pagination
func (s *storage) LoadRange(limit, offset uint) []interface{} {
	if int(offset) >= len(s.items) {
		return []interface{}{}
	}
	//prepare keys for  stable sort
	keys := make([]string, len(s.items))
	i := 0
	for key := range s.items {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	//get items
	items := []interface{}{}
	end := int(offset + limit)
	for i = int(offset); i < end && i < len(keys); i++ {
		value, ok := s.Load(keys[i])
		//if can't get - try next
		if ok {
			items = append(items, value)
		} else {
			end++
		}
	}

	return items
}

//Store sets the value for a key.
func (s *storage) Store(key string, value interface{}) {
	s.mutex.Lock()
	s.items[key] = value
	s.mutex.Unlock()
}

//Delete deletes value from store
func (s *storage) Delete(key string) {
	s.mutex.Lock()
	delete(s.items, key)
	s.mutex.Unlock()
}
