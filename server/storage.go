package server

import (
	"sort"
	"sync"
)

//Storage is concurent safe storage for any struct
type Storage struct {
	mutex sync.RWMutex
	items map[string]interface{}
}

var instance *Storage
var once sync.Once

//GetStorage - get storage as singleton
func GetStorage() *Storage {
	once.Do(func() {
		instance = &Storage{
			mutex: sync.RWMutex{},
			items: map[string]interface{}{},
		}
	})

	return instance
}

//Load returns value stored in the map for a key, or nil if no value present
//The ok result indicates whether value was found in the map.
func (s *Storage) Load(key string) (interface{}, bool) {
	s.mutex.Lock()
	value, ok := s.items[key]
	s.mutex.Unlock()
	return value, ok
}

//LoadRange returns slice with values stored in the map in order with offset and limit
//This is not a production ready code, just for simplifying pagination
func (s *Storage) LoadRange(limit, offset int) []interface{} {
	if int(offset) >= len(s.items) {
		return []interface{}{}
	}
	//prepare keys for  stable sort
	keys := make([]string, len(s.items))
	i := 0
	s.mutex.Lock()
	for key := range s.items {
		keys[i] = key
		i++
	}
	s.mutex.Unlock()
	sort.Strings(keys)
	//get items
	items := []interface{}{}
	end := offset + limit
	for i = offset; i < end && i < len(keys); i++ {
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
func (s *Storage) Store(key string, value interface{}) {
	s.mutex.Lock()
	s.items[key] = value
	s.mutex.Unlock()
}

//Delete deletes value from store
func (s *Storage) Delete(key string) {
	s.mutex.Lock()
	delete(s.items, key)
	s.mutex.Unlock()
}
