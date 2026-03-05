package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"math/rand"
)

// sync.Cond实现多生产者多消费者
func producerAndConsumer() {
	var (
		wg   sync.WaitGroup
		cond sync.Cond
	)
	cond.L = &sync.Mutex{}
	msgChan := make(chan int, 5)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rand.Seed(time.Now().UnixNano())

	// 生产者
	producer := func(ctx context.Context, idx int) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				cond.Broadcast()
				fmt.Printf("producer %d done\n", idx)
				return
			default:
				cond.L.Lock()
				for len(msgChan) == 5 {
					cond.Wait()
				}

				num := rand.Intn(100)
				msgChan <- num
				fmt.Printf("producer %d produced: %d\n", idx, num)
				cond.Signal()
				cond.L.Unlock()
			}
		}
	}

	// 消费者
	consumer := func(ctx context.Context, idx int) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				for len(msgChan) > 0 {
					select {
					case num := <-msgChan:
						fmt.Printf("consumer %d consumed: %d\n", idx, num)
					default:
						break
					}
				}
				fmt.Printf("consumer %d done\n", idx)
				return
			default:
				cond.L.Lock()
				for len(msgChan) == 0 {
					cond.Wait()
				}

				num := <-msgChan
				fmt.Printf("consumer %d consumed: %d\n", idx, num)
				cond.Signal()
				cond.L.Unlock()
			}
		}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go producer(ctx, i)
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go consumer(ctx, i)
	}

	wg.Wait()
	close(msgChan)
	fmt.Println("done")
}
