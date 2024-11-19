package cache

type CacheProvider interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

type Cache struct {
	provider CacheProvider
}

func (*Cache) Construct() interface{} {
	return func() *Cache {
		return &Cache{
			provider: createCacheProvider(),
		}
	}
}

func (c *Cache) Has(key string) bool {
	_, ok := c.provider.Get(key)
	return ok
}

func (c *Cache) Get(key string) (interface{}, bool) {
	return c.provider.Get(key)
}

func (c *Cache) Set(key string, value interface{}) {
	c.provider.Set(key, value)
}

type SimpleCache struct{}

func createCacheProvider() CacheProvider {
	return &SimpleCache{}
}

func (*SimpleCache) Get(key string) (interface{}, bool) {
	return nil, false
}

func (*SimpleCache) Set(key string, value interface{}) {}
