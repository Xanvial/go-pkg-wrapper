package inmemcache

import "errors"

type Config struct {
	MaxSizeCount int // for ccache
	MaxSizeByte  int // for freecache
}

var errEmptyCache = errors.New("empty data")
