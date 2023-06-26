package factory

// 工厂方法模式
// 思想: 由抽象产品/抽象创建工厂/具体产品工厂/具体产品组成,具体的产品实现抽象产品,具体的产品工厂实现抽象工厂, 具体的产品工厂创建具体的产品
// 缺点: 增加了代码复杂度
// 优点: 高度解耦, 容易扩展,新增产品时只需要添加对应的具体产品,具体产品工厂

// Cache -----------抽象产品 Cache------------
type Cache interface {
	Set(key, value string)
	Get(key string) string
}

// RedisCache -----------具体产品 RedisCache--------
type RedisCache struct {
	data map[string]string
}

// Set 添加缓存
func (r *RedisCache) Set(key, value string) {
	r.data[key] = value
}

// Get 获取缓存值
func (r *RedisCache) Get(key string) string {
	return r.data[key]
}

// MemCache ---------------具体产品 MemCache ------------
type MemCache struct {
	data map[string]string
}

// Set 添加缓存
func (m *MemCache) Set(key, value string) {
	m.data[key] = value
}

// Get 获取缓存值
func (m *MemCache) Get(key string) string {
	return m.data[key]
}

// CacheFactory ------------------- 抽象工厂 -------------
type CacheFactory interface {
	Create() Cache
}

// RedisCacheFactory ------------------- 具体产品redis的工厂------------
type RedisCacheFactory struct{}

// Create 创建 RedisCache 实例
func (rf *RedisCacheFactory) Create() Cache {
	return &RedisCache{
		data: make(map[string]string),
	}
}

// MemCacheFactory -------------------- 具体产品 memcache 工厂 -----------
type MemCacheFactory struct{}

// Create 创建 MemCache 实例
func (mf *MemCacheFactory) Create() Cache {
	return &MemCache{
		data: make(map[string]string),
	}
}
