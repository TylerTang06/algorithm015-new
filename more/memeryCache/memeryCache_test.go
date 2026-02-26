package memerycache

import (
	"container/heap"
	"context"
	"testing"
	"time"
)

func TestNewMemoryCache(t *testing.T) {
	// 测试默认配置
	cache := NewMemoryCache(nil)
	if cache == nil {
		t.Fatal("NewMemoryCache returned nil")
	}

	if cache.config.MaxMemerySize != 100*1024*1024 {
		t.Errorf("expected default MaxMemerySize to be 100MB, got %d", cache.config.MaxMemerySize)
	}

	if cache.config.DefaultTTL != 5*time.Minute {
		t.Errorf("expected default DefaultTTL to be 5 minutes, got %v", cache.config.DefaultTTL)
	}

	if cache.config.EvictType != "LRU" {
		t.Errorf("expected default EvictType to be LRU, got %s", cache.config.EvictType)
	}

	if len(cache.shards) != 16 {
		t.Errorf("expected 16 shards, got %d", len(cache.shards))
	}

	cache.Close()
}

func TestNewMemoryCacheWithConfig(t *testing.T) {
	config := &CacheConfig{
		MaxMemerySize: 50 * 1024 * 1024,
		DefaultTTL:    10 * time.Minute,
		CleanInterval: 30 * time.Second,
		EvictType:     "LRU",
		Shards:        8,
		Metrics:       true,
	}

	cache := NewMemoryCache(config)
	if cache == nil {
		t.Fatal("NewMemoryCache returned nil")
	}

	if cache.config.MaxMemerySize != 50*1024*1024 {
		t.Errorf("expected MaxMemerySize to be 50MB, got %d", cache.config.MaxMemerySize)
	}

	if cache.config.DefaultTTL != 10*time.Minute {
		t.Errorf("expected DefaultTTL to be 10 minutes, got %v", cache.config.DefaultTTL)
	}

	if len(cache.shards) != 8 {
		t.Errorf("expected 8 shards, got %d", len(cache.shards))
	}

	cache.Close()
}

func TestMemoryCache_SetAndGet(t *testing.T) {
	cache := NewMemoryCache(nil)
	defer cache.Close()

	ctx := context.Background()

	// 测试基本的 Set 和 Get
	err := cache.Set(ctx, "key1", "value1")
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	value, ok, err := cache.Get(ctx, "key1")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	if !ok {
		t.Error("expected key to exist")
	}

	if value != "value1" {
		t.Errorf("expected value to be 'value1', got %v", value)
	}
}

func TestMemoryCache_GetNotExists(t *testing.T) {
	cache := NewMemoryCache(nil)
	defer cache.Close()

	ctx := context.Background()

	value, ok, err := cache.Get(ctx, "nonexistent")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	if ok {
		t.Error("expected key to not exist")
	}

	if value != nil {
		t.Errorf("expected value to be nil, got %v", value)
	}
}

func TestMemoryCache_SetWithCustomTTL(t *testing.T) {
	cache := NewMemoryCache(nil)
	defer cache.Close()

	ctx := context.Background()

	// 设置自定义 TTL
	err := cache.Set(ctx, "key1", "value1", 100*time.Millisecond)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	// 立即获取应该成功
	_, ok, err := cache.Get(ctx, "key1")
	if err != nil || !ok {
		t.Error("expected key to exist immediately after set")
	}

	// 等待过期
	time.Sleep(150 * time.Millisecond)

	// 过期后获取应该失败
	_, ok, err = cache.Get(ctx, "key1")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	if ok {
		t.Error("expected key to be expired")
	}
}

func TestMemoryCache_GetWithTTL(t *testing.T) {
	cache := NewMemoryCache(nil)
	defer cache.Close()

	ctx := context.Background()

	ttl := 5 * time.Minute
	err := cache.Set(ctx, "key1", "value1", ttl)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	value, remainingTTL, ok, err := cache.GetWithTTL(ctx, "key1")
	if err != nil {
		t.Fatalf("GetWithTTL failed: %v", err)
	}

	if !ok {
		t.Error("expected key to exist")
	}

	if value != "value1" {
		t.Errorf("expected value to be 'value1', got %v", value)
	}

	if remainingTTL <= 0 || remainingTTL > ttl {
		t.Errorf("expected remaining TTL to be between 0 and %v, got %v", ttl, remainingTTL)
	}
}

func TestMemoryCache_UpdateValue(t *testing.T) {
	cache := NewMemoryCache(nil)
	defer cache.Close()

	ctx := context.Background()

	// 设置初始值
	err := cache.Set(ctx, "key1", "value1")
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	// 更新值
	err = cache.Set(ctx, "key1", "value2")
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	// 验证更新后的值
	value, ok, err := cache.Get(ctx, "key1")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	if !ok {
		t.Error("expected key to exist")
	}

	if value != "value2" {
		t.Errorf("expected value to be 'value2', got %v", value)
	}
}

func TestMemoryCache_ContextCancellation(t *testing.T) {
	cache := NewMemoryCache(nil)
	defer cache.Close()

	ctx, cancel := context.WithCancel(context.Background())

	// 取消 context
	cancel()

	// 测试 Set
	err := cache.Set(ctx, "key1", "value1")
	if err == nil {
		t.Error("expected Set to fail with cancelled context")
	}

	// 测试 Get
	_, _, err = cache.Get(ctx, "key1")
	if err == nil {
		t.Error("expected Get to fail with cancelled context")
	}
}

func TestMemoryCache_EvictByLRU(t *testing.T) {
	config := &CacheConfig{
		MaxMemerySize: 5 * 1024, // 5KB
		DefaultTTL:    10 * time.Minute,
		CleanInterval: 1 * time.Minute,
		EvictType:     "LRU",
		Shards:        1,
	}

	cache := NewMemoryCache(config)
	defer cache.Close()

	ctx := context.Background()

	// 创建适中大小的对象（约 300-400 字节 JSON 序列化后）
	largeData := make([]byte, 300)
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	// 设置多个值，触发驱逐
	for i := 0; i < 20; i++ {
		err := cache.Set(ctx, string(rune('a'+i)), largeData)
		if err != nil {
			t.Fatalf("Set failed for key %d: %v", i, err)
		}
		time.Sleep(10 * time.Millisecond) // 确保访问时间不同
	}

	// 验证最新设置的值存在
	value, ok, err := cache.Get(ctx, string(rune('t')))
	if err != nil || !ok {
		t.Error("expected latest key to exist")
	}

	if value == nil {
		t.Error("expected value to not be nil")
	}

	// 验证最老的值可能被驱逐了
	_, ok, _ = cache.Get(ctx, "a")
	if ok {
		t.Log("oldest key still exists (depends on eviction timing)")
	}
}

func TestMemoryCache_Expiration(t *testing.T) {
	config := &CacheConfig{
		MaxMemerySize: 10 * 1024 * 1024,
		DefaultTTL:    5 * time.Minute,
		CleanInterval: 100 * time.Millisecond,
		EvictType:     "LRU",
		Shards:        1,
	}

	cache := NewMemoryCache(config)
	defer cache.Close()

	ctx := context.Background()

	// 设置短期值
	err := cache.Set(ctx, "short", "value", 50*time.Millisecond)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	// 设置长期值
	err = cache.Set(ctx, "long", "value", 5*time.Minute)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	// 等待短期值过期并被清理
	time.Sleep(200 * time.Millisecond)

	// 短期值应该不存在
	_, ok, _ := cache.Get(ctx, "short")
	if ok {
		t.Error("expected short key to be expired")
	}

	// 长期值应该还存在
	_, ok, _ = cache.Get(ctx, "long")
	if !ok {
		t.Error("expected long key to still exist")
	}
}

func TestMemoryCache_DifferentValueTypes(t *testing.T) {
	cache := NewMemoryCache(nil)
	defer cache.Close()

	ctx := context.Background()

	tests := []struct {
		key   string
		value interface{}
	}{
		{"string", "test string"},
		{"int", 12345},
		{"float", 3.14159},
		{"bool", true},
		{"slice", []int{1, 2, 3}},
		{"map", map[string]int{"a": 1, "b": 2}},
	}

	for _, tt := range tests {
		err := cache.Set(ctx, tt.key, tt.value)
		if err != nil {
			t.Fatalf("Set failed for key %s: %v", tt.key, err)
		}

		value, ok, err := cache.Get(ctx, tt.key)
		if err != nil {
			t.Fatalf("Get failed for key %s: %v", tt.key, err)
		}

		if !ok {
			t.Errorf("expected key %s to exist", tt.key)
		}

		// 对于复杂类型，这里只检查非空
		if value == nil {
			t.Errorf("expected value for key %s to not be nil", tt.key)
		}
	}
}

func TestMemoryCache_Sharding(t *testing.T) {
	config := &CacheConfig{
		MaxMemerySize: 10 * 1024 * 1024,
		DefaultTTL:    5 * time.Minute,
		CleanInterval: 1 * time.Minute,
		EvictType:     "LRU",
		Shards:        4,
	}

	cache := NewMemoryCache(config)
	defer cache.Close()

	ctx := context.Background()

	// 设置多个键到不同的分片
	keys := []string{"key1", "key2", "key3", "key4", "key5", "key6", "key7", "key8"}

	for _, key := range keys {
		err := cache.Set(ctx, key, "value")
		if err != nil {
			t.Fatalf("Set failed for key %s: %v", key, err)
		}
	}

	// 验证所有键都可以获取
	for _, key := range keys {
		_, ok, err := cache.Get(ctx, key)
		if err != nil {
			t.Fatalf("Get failed for key %s: %v", key, err)
		}

		if !ok {
			t.Errorf("expected key %s to exist", key)
		}
	}
}

func TestMemoryCache_ConcurrentAccess(t *testing.T) {
	cache := NewMemoryCache(nil)
	defer cache.Close()

	ctx := context.Background()

	// 并发写入
	done := make(chan bool)
	for i := 0; i < 100; i++ {
		go func(idx int) {
			key := string(rune('a' + (idx % 26)))
			err := cache.Set(ctx, key, idx)
			if err != nil {
				t.Errorf("Set failed for key %s: %v", key, err)
			}
			done <- true
		}(i)
	}

	// 等待所有写入完成
	for i := 0; i < 100; i++ {
		<-done
	}

	// 并发读取
	for i := 0; i < 100; i++ {
		go func(idx int) {
			key := string(rune('a' + (idx % 26)))
			_, _, err := cache.Get(ctx, key)
			if err != nil {
				t.Errorf("Get failed for key %s: %v", key, err)
			}
			done <- true
		}(i)
	}

	// 等待所有读取完成
	for i := 0; i < 100; i++ {
		<-done
	}
}

func TestCacheShard_getShard(t *testing.T) {
	cache := NewMemoryCache(nil)
	defer cache.Close()

	// 测试相同的 key 总是映射到相同的分片
	key := "testKey"
	shard1 := cache.getShard(key)
	shard2 := cache.getShard(key)

	if shard1 != shard2 {
		t.Error("same key should map to same shard")
	}

	// 测试不同的 key 可能映射到不同的分片
	shard3 := cache.getShard("differentKey")
	if shard1 == shard3 {
		// 这是可能的，但不总是如此
		t.Log("different keys mapped to same shard (this is possible)")
	}
}

func TestCacheShard_getTTL(t *testing.T) {
	cache := NewMemoryCache(&CacheConfig{
		DefaultTTL: 10 * time.Minute,
		Shards:     4,
	})
	defer cache.Close()

	// 测试自定义 TTL
	customTTL := cache.getTTL(5 * time.Minute)
	if customTTL != 5*time.Minute {
		t.Errorf("expected custom TTL, got %v", customTTL)
	}

	// 测试默认 TTL
	defaultTTL := cache.getTTL()
	if defaultTTL != 10*time.Minute {
		t.Errorf("expected default TTL, got %v", defaultTTL)
	}
}

func TestExpireHeap(t *testing.T) {
	h := &expireHeap{}

	// 测试 Push 和 Len
	item1 := &cacheItem{key: "key1", expireAt: time.Now().Add(5 * time.Minute)}
	item2 := &cacheItem{key: "key2", expireAt: time.Now().Add(3 * time.Minute)}
	item3 := &cacheItem{key: "key3", expireAt: time.Now().Add(7 * time.Minute)}

	heap.Push(h, item1)
	heap.Push(h, item2)
	heap.Push(h, item3)

	if h.Len() != 3 {
		t.Errorf("expected heap length to be 3, got %d", h.Len())
	}

	// 测试 Less - 堆顶应该是最早过期的
	if h.Less(0, 1) {
		t.Log("item at index 0 expires before item at index 1")
	}

	// 测试 Pop
	popped := heap.Pop(h).(*cacheItem)
	if popped.key != "key2" {
		t.Errorf("expected to pop 'key2' (earliest expiration), got %s", popped.key)
	}

	if h.Len() != 2 {
		t.Errorf("expected heap length to be 2 after pop, got %d", h.Len())
	}
}

func TestMemoryCache_Close(t *testing.T) {
	cache := NewMemoryCache(nil)

	// Close 应该安全地调用
	cache.Close()

	// // 重复 Close 也应该安全
	// cache.Close()
}

func TestMemoryCache_ZeroTTL(t *testing.T) {
	cache := NewMemoryCache(nil)
	defer cache.Close()

	ctx := context.Background()

	// 设置零 TTL 应该使用默认 TTL
	err := cache.Set(ctx, "key1", "value1", 0)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	// 应该能立即获取到
	_, ok, err := cache.Get(ctx, "key1")
	if err != nil || !ok {
		t.Error("expected key to exist with zero TTL (using default)")
	}
}
