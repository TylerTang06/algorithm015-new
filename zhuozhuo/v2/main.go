package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
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
   - 所有“成功参与本次执行”的调用：
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

// 执行组：跟踪一次 job 执行的所有调用者
type execution struct {
	mu      sync.Mutex   // 保证并发
	once    sync.Once    // 确保 job 只执行一次
	count   atomic.Int32 // 参与本次执行的调用数量
	result  error        // job 执行结果
	errCh   chan error   // 用于广播结果到所有等待者
	cleaned bool         // 标记是否已清理
}

var (
	// 全局缓存
	executions = make(map[string]*execution)
	existingMu sync.Mutex
)

func target(
	ctx context.Context,
	id string,
	job func(context.Context) error,
) (count int, err error) {
	// 检查是否已存在正在执行的任务
	existingMu.Lock()
	existing, ok := executions[id]
	if !ok {
		// 创建新的执行组
		existing = &execution{
			errCh: make(chan error, 1),
		}
		executions[id] = existing
	}
	existingMu.Unlock()

	// 注册参与本次执行（使用原子操作）
	existing.count.Add(1)

	// 确保只有一个协程执行 job
	existing.once.Do(func() {
		defer func() {
			// 捕获 panic 并转换为 error
			if r := recover(); r != nil {
				existing.result = fmt.Errorf("panic: %v", r)
				existing.errCh <- existing.result
				close(existing.errCh)
			}
		}()

		// 执行 job
		existing.result = job(ctx)
		existing.errCh <- existing.result
		close(existing.errCh)
	})

	// 等待 job 执行完成或 ctx 取消
	select {
	case err := <-existing.errCh:
		count = int(existing.count.Load())

		// 只由第一个完成的协程清理
		existing.mu.Lock()
		if !existing.cleaned {
			existingMu.Lock()
			delete(executions, id)
			existingMu.Unlock()
			existing.cleaned = true
		}
		existing.mu.Unlock()

		return count, err
	case <-ctx.Done():
		// ctx 取消，不计入 count（使用原子操作）
		existing.count.Add(-1)
		return int(existing.count.Load()), ctx.Err()
	}
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
