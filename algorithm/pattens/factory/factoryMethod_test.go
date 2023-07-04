package factory

import (
	"fmt"
	"testing"
)

func TestRedisCache(t *testing.T) {
	var redisCacheFactory CacheFactory = &RedisCacheFactory{}
	redisCache := redisCacheFactory.Create()
	redisCache.Set("uid", "xxxx123213")
	fmt.Println("redisCache[uid] is ", redisCache.Get("uid"))

}
