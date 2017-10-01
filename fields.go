package cinder

import "sort"

// Fields ...
type Fields map[string]interface{}

// Get returns a field value by name.
func (f Fields) Get(k string) interface{} {
	return f[k]
}

// Set changes the value of a key already contained in the Fields object.
func (f Fields) Set(k string, v interface{}) {
	f[k] = v
}

// Keys returns the keys currently in the Fields map.
func (f Fields) Keys() (keys []string) {
	for k := range f {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return
}
