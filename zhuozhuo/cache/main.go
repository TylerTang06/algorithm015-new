package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/**
你需要实现的目标函数 target

@param ctx 控制当前调用的生命周期；
@param id 是一个随机字符串，例如 6A10A467-2842-A460-5353-DBE7D41986B7；
@param job 是一个耗时操作，可能：
        - 正常返回
        - 返回 error
        - panic
@return count 表示【成功参与本次 job 执行】的相同 id 调用数量
@return err   表示 job 的执行结果（如失败）

关键特性说明：

1. 相同 id 并发调用 target：
   - job 只会被执行一次
   - 所有"成功参与本次执行"的调用：
     - 返回相同的 count
     - 返回相同的 err
2. 不同 id：
   - 互不影响，可以并行执行
3. ctx 被取消：
   - 当前调用应立即返回
   - 不应影响其他调用
   - 不应计入 count
4. job 返回 error 或 panic：
   - 所有仍在等待的调用返回 error
5. 相同 id 串行调用：
   - 每次都应独立执行 job
   - 每次返回 count = 1

注意：
- job 的执行时间不固定
- 不允许依赖 sleep 或时间假设
*/

// jobResult 存储 job 执行的最终结果
type jobResult struct {
	count    int       // 参与本次执行的调用数量
	err      error     // job 执行结果
	finished time.Time // 完成时间
	waiters  int       // 正在等待结果的调用者数量
}

// cacheEntry 缓存条目，可以是执行状态或结果
type cacheEntry interface{}

// jobExecution 一次 job 执行的状态
type jobExecution struct {
	mu         sync.Mutex
	registered int            // 已注册的调用数量
	resultCh   chan jobResult // 结果通道
	completed  bool           // 是否已完成
	result     jobResult      // 缓存的结果
	once       sync.Once      // 确保 job 只执行一次
}

var (
	cache   = make(map[string]cacheEntry)
	cacheMu sync.Mutex
)

func target(
	ctx context.Context,
	id string,
	job func(context.Context) error,
) (count int, err error) {
	// 检查缓存
	cacheMu.Lock()
	entry, exists := cache[id]
	if !exists {
		// 创建新的执行状态
		exec := &jobExecution{
			resultCh:   make(chan jobResult, 1),
			completed:  false,
			registered: 0,
		}
		cache[id] = exec
		entry = exec
	}
	cacheMu.Unlock()

	switch v := entry.(type) {
	case *jobResult:
		// 已有结果，需要先增加 waiters 计数
		cacheMu.Lock()
		v.waiters++
		cacheMu.Unlock()

		// 返回后减少等待者计数
		defer func() {
			cacheMu.Lock()
			v.waiters--
			if v.waiters <= 0 {
				delete(cache, id)
			}
			cacheMu.Unlock()
		}()

		return v.count, v.err
	case *jobExecution:
		// 注册
		v.mu.Lock()
		v.registered++
		v.mu.Unlock()

		// 使用 sync.Once 确保 job 只执行一次
		var execResult jobResult
		var resultReady = false

		v.once.Do(func() {
			// 执行 job
			var jobErr error
			var panicErr interface{}

			func() {
				defer func() {
					if r := recover(); r != nil {
						panicErr = r
					}
				}()
				jobErr = job(ctx)
			}()

			v.mu.Lock()
			finalCount := v.registered
			result := jobResult{
				count:    finalCount,
				err:      jobErr,
				finished: time.Now(),
				waiters:  finalCount - 1,
			}
			if panicErr != nil {
				result.err = fmt.Errorf("panic: %v", panicErr)
			}
			v.completed = true
			v.result = result
			v.mu.Unlock()

			// 发送结果给所有等待者
			v.resultCh <- result
			close(v.resultCh)

			// 将结果存入缓存
			cacheMu.Lock()
			cache[id] = &result
			cacheMu.Unlock()

			execResult = result
			resultReady = true

			// 如果没有等待者（串行调用场景），执行者需要清除缓存
			if result.waiters <= 0 {
				cacheMu.Lock()
				delete(cache, id)
				cacheMu.Unlock()
			}
		})

		if resultReady {
			// 执行者：先增加 waiters（模拟自己也是一个等待者）
			cacheMu.Lock()
			execResult.waiters++
			cacheMu.Unlock()

			// 延迟清除缓存，让等待者先读取
			defer func() {
				cacheMu.Lock()
				execResult.waiters--
				if execResult.waiters <= 0 {
					delete(cache, id)
				}
				cacheMu.Unlock()
			}()

			return execResult.count, execResult.err
		}

		// 等待执行完成
		<-v.resultCh

		// 从缓存读取结果
		cacheMu.Lock()
		entry := cache[id]
		cacheMu.Unlock()

		if result, ok := entry.(*jobResult); ok {
			// 返回后减少等待者计数
			defer func() {
				cacheMu.Lock()
				result.waiters--
				if result.waiters <= 0 {
					delete(cache, id)
				}
				cacheMu.Unlock()
			}()

			return result.count, result.err
		}

		return 0, fmt.Errorf("unexpected cache entry type")
	}

	return 0, fmt.Errorf("unexpected cache entry type")
}

//////////////////////////////////////////////
///////// 接下来的代码为测试代码，请勿修改 /////////
//////////////////////////////////////////////

// 用来模拟 job 执行次数
// 不要修改
var (
	counter     int
	counterLock sync.Mutex
)

// 模拟 job 的耗时（不固定）
// 不要修改
const (
	mockJobTimeout = 300 * time.Millisecond
	tolerate       = 30 * time.Millisecond
)

// mock job：计数 + 延时
// 不要修改
func mockJob(ctx context.Context) error {
	select {
	case <-time.After(mockJobTimeout):
	case <-ctx.Done():
		return ctx.Err()
	}

	counterLock.Lock()
	counter++
	counterLock.Unlock()
	return nil
}

// 相同 id 并发调用
// 不要修改
func testCaseSampleIdParallel() {
	counter = 0
	const (
		id     = "CBD225E1-B7D9-BE76-9735-1D0A9B62EE4D"
		repeat = 5
	)

	wg := sync.WaitGroup{}
	wg.Add(repeat)

	tStart := time.Now()
	for i := 0; i < repeat; i++ {
		go func() {
			ctx := context.Background()
			count, err := target(ctx, id, mockJob)
			wg.Done()
			if err != nil {
				panic(err)
			}
			if count != repeat {
				panic(fmt.Sprintln("[parallel] count:", count, "!= repeat:", repeat))
			}
		}()
	}

	wg.Wait()

	if counter != 1 {
		panic(fmt.Sprintln("[parallel] counter:", counter, "!= 1"))
	}

	if time.Since(tStart) > mockJobTimeout+tolerate {
		panic("[parallel] timeout")
	}
}

// 相同 id 串行调用
// 不要修改
func testCaseSampleIdSerial() {
	counter = 0
	const (
		id     = "3E5A5C8D-B254-383B-4F33-F6927578FD11"
		repeat = 2
	)

	tStart := time.Now()
	for i := 0; i < repeat; i++ {
		ctx := context.Background()
		count, err := target(ctx, id, mockJob)
		if err != nil {
			panic(err)
		}
		if count != 1 {
			panic(fmt.Sprintln("[serial] count:", count, "!= 1"))
		}
	}

	if counter != repeat {
		panic(fmt.Sprintln("[serial] counter:", counter, "!= repeat:", repeat))
	}

	if time.Since(tStart) > time.Duration(repeat)*mockJobTimeout+tolerate {
		panic("[serial] timeout")
	}
}

// 不同 id 并发调用
// 不要修改
func testCaseRandomId() {
	counter = 0

	ids := []string{
		"id-3", "id-3", "id-3",
		"id-2", "id-2",
		"id-1",
	}

	wg := sync.WaitGroup{}
	wg.Add(len(ids))

	tStart := time.Now()
	for _, id := range ids {
		id := id
		go func() {
			ctx := context.Background()
			count, err := target(ctx, id, mockJob)
			wg.Done()
			if err != nil {
				panic(err)
			}

			expected := map[string]int{
				"id-1": 1,
				"id-2": 2,
				"id-3": 3,
			}[id]

			if count != expected {
				panic(fmt.Sprintln("[random] id:", id, "count:", count, "!= expected:", expected))
			}
		}()
	}

	wg.Wait()

	if counter != 3 {
		panic(fmt.Sprintln("[random] counter:", counter, "!= 3"))
	}

	if time.Since(tStart) > 3*mockJobTimeout+tolerate {
		panic("[random] timeout")
	}
}

// 不要修改
func main() {
	// 先测试串行调用
	testCaseSampleIdSerial()
	fmt.Println("Serial test passed!")

	// 再测试并发
	testCaseSampleIdParallel()
	fmt.Println("Parallel test passed!")

	// 测试不同 ID
	testCaseRandomId()
	fmt.Println("Random ID test passed!")

	const repeat = 50
	for i := 0; i < repeat; i++ {
		testCaseSampleIdParallel()
		testCaseSampleIdSerial()
		testCaseRandomId()
		fmt.Print("\r", i+1, "/", repeat, " ✔ ")
	}
	fmt.Println("\n🎉 All Tests Passed!")
}
