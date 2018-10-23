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

// DictVal the value mapped to DictKey of the dictionary
type DictVal string

// Dictionary the set of elements
type Dictionary struct {
	elements map[DictKey]DictVal
	lock     sync.RWMutex
}

// Put inserts a new element to the dictionary
func (dict *Dictionary) Put(key DictKey, value DictVal) {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	if dict.elements == nil {
		dict.elements = make(map[DictKey]DictVal)
	}
	dict.elements[key] = value
}

// Remove removes a value from the dictionary, given its key
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

// Contains returns true if the key exists in the dictionary
func (dict *Dictionary) Contains(key DictKey) bool {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	var exists bool
	_, exists = dict.elements[key]
	return exists
}

// Find returns the value associated with the key
func (dict *Dictionary) Find(key DictKey) DictVal {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	return dict.elements[key]
}

// Reset removes all the items from the dictionary
func (dict *Dictionary) Reset() {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	dict.elements = make(map[DictKey]DictVal)
}

// NumberOfElements returns the number of elements in the dictionary
func (dict *Dictionary) NumberOfElements() int {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	return len(dict.elements)
}

// GetKeys returns a slice of all the dictionary keys present
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

// Values returns a slice of all the values present
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
