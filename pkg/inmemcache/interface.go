package inmemcache

import "time"

// TODO: WIP, not optimized at all, need to find more general way for multiple libraries
type InMemCache[T any] interface {
	Get(key string) (T, error)
	Set(key string, val T, dur time.Duration)
	// Fetch will try to get the data from cache,
	// but if it's empty or expire will execute fetch function and populate the cache
	Fetch(key string, dur time.Duration, fetch func() (T, error)) (T, error)
	Delete(key string)
}
