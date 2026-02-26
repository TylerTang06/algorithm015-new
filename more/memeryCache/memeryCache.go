package memerycache

import (
	"container/heap"
	"container/list"
	"context"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"sync"
	"time"
)

// CacheConfig 缓存配置
type CacheConfig struct {
	// MaxItems   int           // 最大缓存数量
	MaxMemerySize int           // 最大内存大小
	DefaultTTL    time.Duration // 默认过期时间
	CleanInterval time.Duration // 清理间隔
	EvictType     string        // 驱逐类型:LRU/LFU/FIFO/RANDOM等(应该使用枚举类型),默认LRU
	Logger        Logger        // 日志接口
	Shards        int           // 分片数量,提高并发能力
	Compress      bool          // 是否压缩
	Metrics       bool          // 是否启用指标
}

// cacheStats 缓存统计信息
type cacheStats struct {
	Hits          int // 命中次数
	Misses        int // 未命中次数
	Evicts        int // 驱逐次数
	Size          int // 当前缓存大小
	ItemsCount    int // 项目计数
	MaxMemerySize int // 最大内存大小
	Expires       int // 过期次数
	GetSuccess    int // 获取成功次数
	GetFailed     int // 获取失败次数
	SetSuccess    int // 设置成功次数
	SetFailed     int // 设置失败次数
	DeleteSuccess int //	删除成功次数
	DeleteFailed  int // 删除失败次数
	ClearOps      int // 清空操作次数
	// MaxItems   int // 最大缓存数量
}

// cacheItem 缓存项
type cacheItem struct {
	key          string        // key
	value        interface{}   // 值
	expireAt     time.Time     // 过期时间
	accessedAt   time.Time     // 访问时间
	size         int           // 大小
	accessedSize int           // 访问大小
	element      *list.Element // 链表元素
	index        int           // 索引
	// priority  int           // 优先级
}

// cacheShard 缓存分片
type cacheShard struct {
	mu            sync.RWMutex          // 读写锁
	items         map[string]*cacheItem // 缓存项
	stats         *cacheStats           // 统计信息
	logger        Logger                // 日志接口
	evictType     string                //	驱逐类型
	size          int                   // 当前缓存大小
	maxMemerySize int                   // 最大内存大小
	itemsCount    int                   // 项目计数
	list          *list.List            // 链表
	heap          *expireHeap           // 过期堆
	// maxItems   int                   // 最大缓存数量
}

// expireHeap 过期堆
type expireHeap []*cacheItem

func (h expireHeap) Len() int { return len(h) }

func (h expireHeap) Less(i, j int) bool {
	return h[i].expireAt.Before(h[j].expireAt)
}

func (h expireHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *expireHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*cacheItem)
	item.index = n
	*h = append(*h, item)
}

func (h *expireHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	item.index = -1
	*h = old[0 : n-1]
	return item
}

// memoryCache 内存缓存
type memoryCache struct {
	config  *CacheConfig
	stats   *cacheStats
	shards  []*cacheShard
	logger  Logger
	closeCh chan struct{}
	wg      sync.WaitGroup
}

// NewmemoryCache 创建一个新的内存缓存
func NewmemoryCache(config *CacheConfig) *memoryCache {
	if config == nil {
		config = &CacheConfig{
			// MaxItems:   10000,
			MaxMemerySize: 100 * 1024 * 1024, // 100MB
			DefaultTTL:    5 * time.Minute,
			CleanInterval: 1 * time.Minute,
			EvictType:     "LRU",
			Logger:        &defaultLogger{},
			Shards:        16,
			Compress:      false,
			Metrics:       true,
		}
	}

	if config.Logger == nil {
		config.Logger = &defaultLogger{}
	}

	cache := &memoryCache{
		config: config,
		stats: &cacheStats{
			MaxMemerySize: config.MaxMemerySize,
		},
		shards:  make([]*cacheShard, config.Shards),
		logger:  config.Logger,
		closeCh: make(chan struct{}),
	}

	// 初始化 shards
	maxMemerySize := config.MaxMemerySize / config.Shards
	for i := 0; i < config.Shards; i++ {

		cache.shards[i] = &cacheShard{
			items: map[string]*cacheItem{},
			stats: &cacheStats{
				MaxMemerySize: maxMemerySize,
			},
			logger:        config.Logger,
			evictType:     config.EvictType,
			list:          list.New(),
			heap:          &expireHeap{},
			maxMemerySize: maxMemerySize,
		}
	}

	// 启动定时任务
	if config.CleanInterval > 0 {
		cache.wg.Add(1)
		go cache.cleanupLoop()
	}

	return cache
}

// getShard 获取对应的分片
func (c *memoryCache) getShard(key string) *cacheShard {
	// 使用简单的FNV哈希
	h := fnv.New32a()
	h.Write([]byte(key))
	hash := h.Sum32()
	return c.shards[hash%uint32(len(c.shards))]
}

// cleanupLoop 定期清理过期数据
func (c *memoryCache) cleanupLoop() {
	defer c.wg.Done()

	ticker := time.NewTicker(c.config.CleanInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.cleanupExpired()
		case <-c.closeCh:
			c.logger.Infof("cleanupLoop stopped")
			return
		}
	}
}

// cleanupExpired 清理过期数据
func (c *memoryCache) cleanupExpired() {
	now := time.Now()
	expiredCount := 0

	for _, shard := range c.shards {
		shard.mu.Lock()

		for shard.heap.Len() > 0 {
			item := (*shard.heap)[0]
			if item.expireAt.After(now) {
				break
			}

			// 从堆和map中删除过期条目
			heap.Pop(shard.heap)
			shard.removeFromList(item)
			delete(shard.items, item.key)
			shard.size -= item.size
			shard.stats.Expires++
			shard.stats.ItemsCount--
			shard.stats.Size = shard.size
			expiredCount++
		}

		shard.mu.Unlock()
	}

	if expiredCount > 0 {
		c.logger.Infof("cleanupExpired %d expired entries", expiredCount)
	}
}

// Set 设置缓存
func (c *memoryCache) Set(ctx context.Context, key string, value interface{}, ttl ...time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	shard := c.getShard(key)
	return shard.set(key, value, c.getTTL(ttl...))
}

// getTTL 获取TTL
func (c *memoryCache) getTTL(ttl ...time.Duration) time.Duration {
	if len(ttl) > 0 && ttl[0] > 0 {
		return ttl[0]
	}
	return c.config.DefaultTTL
}

// Get 获取缓存
func (c *memoryCache) Get(ctx context.Context, key string) (interface{}, bool, error) {
	select {
	case <-ctx.Done():
		return nil, false, ctx.Err()
	default:
	}

	shard := c.getShard(key)
	return shard.get(key)
}

// GetWithTTL 获取缓存和TTL
func (c *memoryCache) GetWithTTL(ctx context.Context, key string) (interface{}, time.Duration, bool, error) {
	select {
	case <-ctx.Done():
		return nil, 0, false, ctx.Err()
	default:
	}

	shard := c.getShard(key)
	return shard.getWithTTL(key)
}

// removeFromList 从列表中移除条目
func (s *cacheShard) removeFromList(item *cacheItem) {
	if item.element != nil {
		s.list.Remove(item.element)
		item.element = nil
	}
}

// set 设置缓存
func (s *cacheShard) set(key string, value interface{}, ttl time.Duration) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 序列化计算大小
	data, err := json.Marshal(value)
	if err != nil {
		s.stats.SetFailed++
		s.logger.Errorf("Failed to marshal value: %v", err)
		return err
	}

	size := len(data)

	// 检查容量
	if s.size+size > s.maxMemerySize {
		if !s.evict(size) {
			s.stats.SetFailed++
			return fmt.Errorf("insufficient cache capacity")
		}
	}

	now := time.Now()
	expireAt := now.Add(ttl)

	entry := &cacheItem{
		key:          key,
		value:        value,
		size:         size,
		expireAt:     expireAt,
		accessedAt:   now,
		accessedSize: 1,
	}

	// 添加或更新
	if existing, ok := s.items[key]; ok {
		// 更新现有项
		s.size -= existing.size
		s.removeFromHeap(existing)
		s.removeFromList(existing)

		*existing = *entry
	} else {
		// 添加新项
		s.items[key] = entry
		s.stats.ItemsCount++
	}

	// 添加数据结构
	s.addToList(entry)
	s.addToHeap(entry)
	s.size += size

	s.stats.Size = s.size
	s.stats.SetSuccess++

	return nil
}

// addToList 添加到列表
func (s *cacheShard) addToList(item *cacheItem) {
	if s.evictType == "LRU" {
		item.element = s.list.PushFront(item)
	}
}

// addToHeap 添加到堆
func (s *cacheShard) addToHeap(item *cacheItem) {
	heap.Push(s.heap, item)
}

// removeFromHeap 从堆中移除条目
func (s *cacheShard) removeFromHeap(item *cacheItem) {
	if item.index >= 0 {
		heap.Remove(s.heap, item.index)
	}
}

// evict 淘汰
func (s *cacheShard) evict(needSize int) bool {
	for s.size+needSize > s.maxMemerySize && len(s.items) > 0 {
		var item *cacheItem

		switch s.evictType {
		case "LRU":
			// 淘汰最久未使用
			item = s.getEvictItemByLRU()
		case "LFU":
			// TODO:
		case "FIFO":
			// TODO:
		case "RANDOM":
			// TODO:
		default:
			// 默认使用LRU
			item = s.getEvictItemByLRU()
		}

		if item != nil {
			s.removeItem(item)
			s.stats.Evicts++
			s.logger.Debugf("Evicted key: %s, size: %d, policy: %s",
				item.key, item.size, s.evictType)
		} else {
			break
		}
	}

	return s.size+needSize <= s.maxMemerySize
}

// getEvictItemByLRU 获取淘汰的LRU条目
func (s *cacheShard) getEvictItemByLRU() (item *cacheItem) {
	if elem := s.list.Back(); elem != nil {
		// s.list.Remove(elem)
		item = elem.Value.(*cacheItem)
	}
	return
}

// removeItem 移除条目
func (s *cacheShard) removeItem(item *cacheItem) {
	s.removeFromList(item)
	s.removeFromHeap(item)
	delete(s.items, item.key)
	s.size -= item.size
	s.stats.ItemsCount--
	s.stats.Size = s.size
}

// get 获取缓存
func (s *cacheShard) get(key string) (interface{}, bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, ok := s.items[key]
	if !ok {
		s.stats.Misses++
		s.stats.GetFailed++
		return nil, false, nil
	}

	// 检查是否过期
	if time.Now().After(item.expireAt) {
		s.removeItem(item)
		s.stats.Expires++
		s.stats.Misses++
		s.stats.GetFailed++
		return nil, false, nil
	}

	// 更新访问信息
	item.accessedAt = time.Now()
	item.accessedSize++

	// 更新LRU链表位置
	if s.evictType == "LRU" && item.element != nil {
		s.list.MoveToFront(item.element)
	}

	s.stats.Hits++
	s.stats.GetSuccess++

	return item.value, true, nil
}

// getWithTTL 获取缓存 并返回ttl
func (s *cacheShard) getWithTTL(key string) (interface{}, time.Duration, bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item, ok := s.items[key]
	if !ok {
		s.stats.Misses++
		s.stats.GetFailed++
		return nil, 0, false, nil
	}

	now := time.Now()
	if now.After(item.expireAt) {
		s.stats.Misses++
		s.stats.GetFailed++
		return nil, 0, false, nil
	}

	ttl := item.expireAt.Sub(now)

	s.stats.Hits++
	s.stats.GetSuccess++

	return item.value, ttl, true, nil
}

// Close 关闭缓存 避免重读调用
func (c *memoryCache) Close() {
	close(c.closeCh)
	c.wg.Wait()
}
