package dictionary

import "sync"

// Key holds the key of a dictionary item.
// It is defined as an interface so it can be a generic type.
type Key interface{}

// Value holds the value of a dictionary item.
// It is defined as an interface so it can be a generic type.
type Value interface{}

// Dictionary is an object with key of type Key and value of type Value.
type Dictionary struct {
	items map[Key]Value
	lock  sync.RWMutex
}

// Add adds a new item to the dictionary.
// If the dictionary's map is nil it will initialize it.
func (d *Dictionary) Add(key Key, value Value) {
	d.lock.Lock()
	defer d.lock.Unlock()

	if d.items == nil {
		d.items = make(map[Key]Value)
	}

	d.items[key] = value
}

// Remove removes a value from the dictionary by key.
func (d *Dictionary) Remove(key Key) bool {
	d.lock.Lock()
	defer d.lock.Unlock()

	_, ok := d.items[key]
	if ok {
		delete(d.items, key)
	}

	return ok
}

// Get gets the value associate with the key.
func (d *Dictionary) Get(key Key) Value {
	d.lock.RLock()
	defer d.lock.RUnlock()

	return d.items[key]
}

// Exists checks if a key exists in the dictionary.
func (d *Dictionary) Exists(key Key) bool {
	d.lock.RLock()
	defer d.lock.RUnlock()

	_, ok := d.items[key]

	return ok
}

// Clear removes all items from the dictionary.
func (d *Dictionary) Clear() {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.items = make(map[Key]Value)
}

// Length returns the number of items in the dictionary.
func (d *Dictionary) Length() int {
	d.lock.RLock()
	defer d.lock.RUnlock()

	return len(d.items)
}

// Keys returns a slice of all dictionary keys.
func (d *Dictionary) Keys() []Key {
	d.lock.RLock()
	d.lock.RUnlock()

	keys := []Key{}
	for i := range d.items {
		keys = append(keys, i)
	}

	return keys
}

// Values returns a slice of all dictionary values.
func (d *Dictionary) Values() []Value {
	d.lock.RLock()
	defer d.lock.RUnlock()

	values := []Value{}

	for i := range d.items {
		values = append(values, d.items[i])
	}

	return values
}
