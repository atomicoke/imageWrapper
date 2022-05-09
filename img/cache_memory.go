package img

func InitMemoryCache() *MemoryCache {
	return &MemoryCache{
		map[string]*Wrap{},
	}
}

func (m *MemoryCache) Put(key string, wrap *Wrap) {
	m.cache[key] = wrap
}

func (m *MemoryCache) Get(key string) (wrap *Wrap, ok bool) {
	wrap, ok = m.cache[key]
	return wrap, ok
}
