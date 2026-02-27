package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// TestTargetParallelSameId 测试相同 ID 并发调用
func TestTargetParallelSameId(t *testing.T) {
	// 重置全局变量
	executions = make(map[string]*execution)
	counter = 0

	const (
		id     = "test-parallel-id"
		repeat = 10
	)

	var wg sync.WaitGroup
	wg.Add(repeat)

	var counts []int
	var mu sync.Mutex

	for i := 0; i < repeat; i++ {
		go func() {
			defer wg.Done()
			ctx := context.Background()
			count, err := target(ctx, id, mockJob)
			
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			
			mu.Lock()
			counts = append(counts, count)
			mu.Unlock()
		}()
	}

	wg.Wait()

	// 验证所有调用返回的 count 都等于 repeat
	for i, count := range counts {
		if count != repeat {
			t.Errorf("call %d: count = %d, want %d", i, count, repeat)
		}
	}

	// 验证 job 只执行了一次
	if counter != 1 {
		t.Errorf("counter = %d, want 1", counter)
	}
}

// TestTargetSerialSameId 测试相同 ID 串行调用
func TestTargetSerialSameId(t *testing.T) {
	executions = make(map[string]*execution)
	counter = 0

	const (
		id     = "test-serial-id"
		repeat = 3
	)

	for i := 0; i < repeat; i++ {
		ctx := context.Background()
		count, err := target(ctx, id, mockJob)
		
		if err != nil {
			t.Errorf("iteration %d: unexpected error: %v", i, err)
			return
		}
		
		if count != 1 {
			t.Errorf("iteration %d: count = %d, want 1", i, count)
		}
	}

	// 串行调用每次都应独立执行 job
	if counter != repeat {
		t.Errorf("counter = %d, want %d", counter, repeat)
	}
}

// TestTargetDifferentIds 测试不同 ID 并发调用
func TestTargetDifferentIds(t *testing.T) {
	executions = make(map[string]*execution)
	counter = 0

	ids := []string{
		"id-1", "id-1", "id-1",
		"id-2", "id-2",
		"id-3",
	}

	expected := map[string]int{
		"id-1": 3,
		"id-2": 2,
		"id-3": 1,
	}

	var wg sync.WaitGroup
	wg.Add(len(ids))

	results := make(map[string][]int)
	var mu sync.Mutex

	for _, id := range ids {
		id := id
		go func() {
			defer wg.Done()
			ctx := context.Background()
			count, err := target(ctx, id, mockJob)
			
			if err != nil {
				t.Errorf("id %s: unexpected error: %v", id, err)
				return
			}
			
			mu.Lock()
			results[id] = append(results[id], count)
			mu.Unlock()
		}()
	}

	wg.Wait()

	// 验证每个 ID 的 count 值
	for id, expectedCount := range expected {
		counts, ok := results[id]
		if !ok {
			t.Errorf("no results for id %s", id)
			continue
		}
		for i, count := range counts {
			if count != expectedCount {
				t.Errorf("id %s call %d: count = %d, want %d", id, i, count, expectedCount)
			}
		}
	}

	// 验证 job 执行了 3 次（3 个不同的 ID）
	if counter != 3 {
		t.Errorf("counter = %d, want 3", counter)
	}
}

// TestTargetContextCancellation 测试 context 取消
func TestTargetContextCancellation(t *testing.T) {
	executions = make(map[string]*execution)
	counter = 0

	const (
		id         = "test-cancel-id"
		repeat     = 5
		cancelTime = 100 * time.Millisecond
	)

	// 创建一个会超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), cancelTime)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(repeat)

	var cancelledCount atomic.Int64
	var completedCount atomic.Int64

	for i := 0; i < repeat; i++ {
		go func() {
			defer wg.Done()
			count, err := target(ctx, id, mockJob)
			
			if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
				cancelledCount.Add(1)
				// 被取消的调用应该返回减去已取消的 count
				return
			}
			
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			
			completedCount.Add(1)
			_ = count
		}()
	}

	wg.Wait()

	// 由于 job 执行时间是 300ms，而 context 在 100ms 后取消，
	// 所有调用都应该被取消
	if cancelledCount.Load() != int64(repeat) {
		t.Errorf("cancelled calls = %d, want %d", cancelledCount.Load(), repeat)
	}

	if completedCount.Load() != 0 {
		t.Errorf("completed calls = %d, want 0", completedCount.Load())
	}
}

// TestTargetJobError 测试 job 返回错误
func TestTargetJobError(t *testing.T) {
	executions = make(map[string]*execution)
	counter = 0

	const (
		id     = "test-error-id"
		repeat = 3
	)

	expectedErr := errors.New("job failed")

	var wg sync.WaitGroup
	wg.Add(repeat)

	var errorCount atomic.Int64

	for i := 0; i < repeat; i++ {
		go func() {
			defer wg.Done()
			ctx := context.Background()
			count, err := target(ctx, id, func(ctx context.Context) error {
				return expectedErr
			})
			
			if err == nil {
				t.Error("expected error, got nil")
				return
			}
			
			if err != expectedErr {
				t.Errorf("err = %v, want %v", err, expectedErr)
				return
			}
			
			if count != repeat {
				t.Errorf("count = %d, want %d", count, repeat)
				return
			}
			
			errorCount.Add(1)
		}()
	}

	wg.Wait()

	if errorCount.Load() != int64(repeat) {
		t.Errorf("error calls = %d, want %d", errorCount.Load(), repeat)
	}
}

// TestTargetJobPanic 测试 job panic
func TestTargetJobPanic(t *testing.T) {
	executions = make(map[string]*execution)
	counter = 0

	const (
		id     = "test-panic-id"
		repeat = 3
	)

	expectedMsg := "test panic"

	var wg sync.WaitGroup
	wg.Add(repeat)

	var recoveredCount atomic.Int64

	for i := 0; i < repeat; i++ {
		go func() {
			defer wg.Done()
			ctx := context.Background()
			count, err := target(ctx, id, func(ctx context.Context) error {
				panic(expectedMsg)
			})
			
			if err == nil {
				t.Error("expected error, got nil")
				return
			}
			
			if err.Error() != fmt.Sprintf("panic: %s", expectedMsg) {
				t.Errorf("err = %v, want panic: %s", err, expectedMsg)
				return
			}
			
			if count != repeat {
				t.Errorf("count = %d, want %d", count, repeat)
				return
			}
			
			recoveredCount.Add(1)
		}()
	}

	wg.Wait()

	if recoveredCount.Load() != int64(repeat) {
		t.Errorf("recovered calls = %d, want %d", recoveredCount.Load(), repeat)
	}
}

// TestTargetSingleCall 测试单次调用
func TestTargetSingleCall(t *testing.T) {
	executions = make(map[string]*execution)
	counter = 0

	ctx := context.Background()
	id := "test-single-id"
	
	count, err := target(ctx, id, mockJob)
	
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	
	if count != 1 {
		t.Errorf("count = %d, want 1", count)
	}
	
	if counter != 1 {
		t.Errorf("counter = %d, want 1", counter)
	}
}

// TestTargetConcurrentDifferentIds 测试大量并发不同 ID
func TestTargetConcurrentDifferentIds(t *testing.T) {
	executions = make(map[string]*execution)
	counter = 0

	const (
		numIds    = 50
		perId     = 3
	)

	var wg sync.WaitGroup
	wg.Add(numIds * perId)

	for i := 0; i < numIds; i++ {
		id := fmt.Sprintf("id-%d", i)
		for j := 0; j < perId; j++ {
			go func(id string) {
				defer wg.Done()
				ctx := context.Background()
				count, err := target(ctx, id, mockJob)
				
				if err != nil {
					t.Errorf("id %s: unexpected error: %v", id, err)
					return
				}
				
				if count != perId {
					t.Errorf("id %s: count = %d, want %d", id, count, perId)
				}
			}(id)
		}
	}

	wg.Wait()

	// 每个 ID 应执行一次 job
	if counter != numIds {
		t.Errorf("counter = %d, want %d", counter, numIds)
	}
}

// TestTargetMixedCancellation 测试混合正常和取消的调用
func TestTargetMixedCancellation(t *testing.T) {
	executions = make(map[string]*execution)
	counter = 0

	id := "test-mixed-id"
	
	// 创建一个会被取消的 context
	ctxCancel, cancel := context.WithCancel(context.Background())
	
	var wg sync.WaitGroup
	
	// 启动几个正常的调用
	normalCount := 3
	wg.Add(normalCount)
	for i := 0; i < normalCount; i++ {
		go func() {
			defer wg.Done()
			ctx := context.Background()
			count, err := target(ctx, id, mockJob)
			
			if err != nil {
				t.Errorf("normal call failed: %v", err)
				return
			}
			
			if count != normalCount {
				t.Errorf("normal call: count = %d, want %d", count, normalCount)
			}
		}()
	}
	
	// 稍后取消
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()
	
	// 启动几个会被取消的调用
	cancelCalls := 2
	wg.Add(cancelCalls)
	for i := 0; i < cancelCalls; i++ {
		go func() {
			defer wg.Done()
			count, err := target(ctxCancel, id, mockJob)
			
			if !errors.Is(err, context.Canceled) {
				t.Errorf("expected Canceled error, got: %v", err)
				return
			}
			
			_ = count
		}()
	}
	
	wg.Wait()
}

// BenchmarkTargetParallel 基准测试：并发性能
func BenchmarkTargetParallel(b *testing.B) {
	executions = make(map[string]*execution)
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		id := fmt.Sprintf("bench-id-%d", i%10) // 限制 ID 数量以测试缓存效果
		
		var wg sync.WaitGroup
		wg.Add(10)
		
		for j := 0; j < 10; j++ {
			go func() {
				defer wg.Done()
				ctx := context.Background()
				_, _ = target(ctx, id, mockJob)
			}()
		}
		
		wg.Wait()
	}
}
