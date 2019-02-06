//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and sync package
import (
	"fmt"
	"sync"
)

// DictKey is the key of the dictionary
type DictKey string

// DictVal type
type DictVal string

// Dictionary class
type Dictionary struct {
	elements map[DictKey]DictVal
	lock     sync.RWMutex
}

// Put method
func (dict *Dictionary) Put(key DictKey, value DictVal) {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	if dict.elements == nil {
		dict.elements = make(map[DictKey]DictVal)
	}
	dict.elements[key] = value
}

// Remove method
func (dict *Dictionary) Remove(key DictKey) bool {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	var exists bool
	_, exists = dict.elements[key]
	if exists {
		delete(dict.elements, key)
	}
	return exists
}

// Contains method
func (dict *Dictionary) Contains(key DictKey) bool {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	var exists bool
	_, exists = dict.elements[key]
	return exists
}

// Find method
func (dict *Dictionary) Find(key DictKey) DictVal {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	return dict.elements[key]
}

// Reset method
func (dict *Dictionary) Reset() {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	dict.elements = make(map[DictKey]DictVal)
}

// NumberOfElements method
func (dict *Dictionary) NumberOfElements() int {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	return len(dict.elements)
}

// GetKeys method
func (dict *Dictionary) GetKeys() []DictKey {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	var dictKeys []DictKey
	dictKeys = []DictKey{}
	var key DictKey
	for key = range dict.elements {
		dictKeys = append(dictKeys, key)
	}
	return dictKeys
}

// GetValues method
func (dict *Dictionary) GetValues() []DictVal {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	var dictValues []DictVal
	dictValues = []DictVal{}
	var key DictKey
	for key = range dict.elements {
		dictValues = append(dictValues, dict.elements[key])
	}
	return dictValues
}

// main method
func main() {

	var dict *Dictionary = &Dictionary{}

	dict.Put("1", "1")
	dict.Put("2", "2")
	dict.Put("3", "3")
	dict.Put("4", "4")

	fmt.Println(dict)

}
