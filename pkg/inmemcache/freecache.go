package inmemcache

import (
	"encoding/json"
	"time"

	"github.com/coocood/freecache"
)

type freeCache[T any] struct {
	cache *freecache.Cache
}

func NewFreeCache[T any](cfg Config) InMemCache[T] {
	return &freeCache[T]{
		cache: freecache.NewCache(cfg.MaxSizeCount),
	}
}

func (fc *freeCache[T]) Get(key string) (T, error) {
	var result T
	data, err := fc.cache.Get([]byte(key))
	if err != nil {
		return result, err
	}

	json.Unmarshal(data, &result)
	return result, nil
}

func (fc *freeCache[T]) Set(key string, val T, dur time.Duration) error {
	data, err := json.Marshal(val)
	if err != nil {
		// do nothing
		return err
	}

	fc.cache.Set([]byte(key), data, int(dur.Seconds()))
	return nil
}

func (fc *freeCache[T]) Fetch(key string, dur time.Duration, fetch func() (T, error)) (T, error) {
	var result T
	tmpData, err := fc.cache.Get([]byte(key))
	if tmpData != nil {
		json.Unmarshal(tmpData, &result)
		return result, nil
	}

	value, err := fetch()
	if err != nil {
		return value, err
	}

	fc.Set(key, value, dur)
	return value, nil
}

func (fc *freeCache[T]) Delete(key string) error {
	fc.cache.Del([]byte(key))
}
