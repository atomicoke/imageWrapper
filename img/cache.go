package img

import log "github.com/sirupsen/logrus"

var c Cache

func InitCache(model CacheModel) {
	switch model {
	case MEMORY:
		c = InitMemoryCache()
	default:
		c = InitMemoryCache()
	}
	log.Infoln("Cache initialized with model: ", model)
}

func PutToCache(key string, wrap *Wrap) {
	log.Debugln("初始化   : " + key)
	c.Put(key, wrap)
}

func GetFromCache(key string) (wrap *Wrap, ok bool) {
	wrap, ok = c.Get(key)
	if ok {
		log.Debugln("命中缓存 : " + key)
	}
	return wrap, ok
}
