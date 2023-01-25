### Sample usage

```
// initialize cache object with specific type, like below example using int
tmpCache := inmemcache.NewCCache[int](inmemcache.Config{})

// Set cache data and duration
tmpCache.Set("data1", 123123, time.Minute)

// Get the data and print it
res := tmpCache.Get("data1")
log.Println("res:", res)
```