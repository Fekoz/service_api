package cache

import (
	"fmt"
	"log"

	conf "bitbucket.org/klaraeng/package_config/conf"
	"github.com/bradfitz/gomemcache/memcache"
)

type Service struct {
	cached *memcache.Client
	conf   *conf.Params
}

func NewCacheService(confCache *conf.Data, params *conf.Params) *Service {
	mc := memcache.New(fmt.Sprintf("%s:%s", confCache.Cache.Addr, confCache.Cache.Port))
	if mc == nil {
		log.Panic("Error up cache to [", confCache.Cache.Addr, ":", confCache.Cache.Port, "]")
	}
	return &Service{
		mc,
		params,
	}
}

func (c *Service) append(key string, value []byte, timeMin int64) {
	//fmt.Println(fmt.Sprintf("[%s]\tappend\t%s", time.Now().String(), key))
	c.cached.Add(&memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: int32(timeMin),
	})
}

func (c *Service) is(key string) bool {
	//fmt.Println(fmt.Sprintf("[%s]\tis\t%s", time.Now().String(), key))
	_, err := c.cached.Get(key)
	if err != nil {
		return false
	}

	return true
}

func (c *Service) read(key string) []byte {
	//fmt.Println(fmt.Sprintf("[%s]\tread\t%s", time.Now().String(), key))
	cache, err := c.cached.Get(key)
	if err != nil {
		return nil
	}
	return cache.Value
}

func (c *Service) drop(key string) {
	//fmt.Println(fmt.Sprintf("[%s]\tdrop\t%s", time.Now().String(), key))
	_ = c.cached.Delete(key)
}

func (c *Service) rewrite(key string, value []byte, timeMin int64) {
	//fmt.Println(fmt.Sprintf("[%s]\trewrite\t%s", time.Now().String(), key))
	_ = c.cached.Replace(&memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: int32(timeMin),
	})
}

func (c *Service) FlushAll() {
	_ = c.cached.FlushAll()
}
