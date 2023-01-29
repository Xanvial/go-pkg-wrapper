package inmemcache

import (
	"time"

	"github.com/karlseguin/ccache/v3"
)

type cCache[T any] struct {
	cache *ccache.Cache[T]
}

func NewCCache[T any](cfg Config) InMemCache[T] {
	configuration := ccache.Configure[T]()
	if cfg.MaxSizeCount > 0 {
		configuration.MaxSize(int64(cfg.MaxSizeCount))
	}
	var cache = ccache.New(configuration)
	return &cCache[T]{
		cache: cache,
	}
}

func (cc *cCache[T]) Get(key string) (T, error) {
	data := cc.cache.Get(key)
	if data == nil || data.Expired() {
		var result T
		return result, errEmptyCache
	}

	return data.Value(), nil
}

func (cc *cCache[T]) Set(key string, val T, dur time.Duration) error {
	cc.cache.Set(key, val, dur)
	return nil
}

func (cc *cCache[T]) Fetch(key string, dur time.Duration, fetch func() (T, error)) (T, error) {
	data, err := cc.cache.Fetch(key, dur, fetch)
	if err != nil {
		var result T
		return result, err
	}

	return data.Value(), nil
}

func (cc *cCache[T]) Delete(key string) error {
	cc.cache.Delete(key)
	return nil
}
