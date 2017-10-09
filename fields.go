package cinder

import "sort"

// Fields ...
type Fields map[string]interface{}

// Keys returns the keys currently in the Fields map. The returned strings are sorted.
func (f Fields) Keys() (keys []string) {
	for k := range f {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return
}
