package main

import (
	"fmt"
	"sync"
)

// 三个goroutinue交替打印abc 10次
func printCh() {
	ch1, ch2, ch3 := make(chan struct{}, 1), make(chan struct{}, 1), make(chan struct{}, 1)

	var f func(ch1, ch2 chan struct{}, s string)
	var wg sync.WaitGroup

	f = func(ch1, ch2 chan struct{}, s string) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ch1
			fmt.Println(s)
			ch2 <- struct{}{}
		}
		// 完成10次循环后不再等待，让 wg.Done() 被调用
	}

	wg.Add(3)
	go f(ch1, ch2, "a")
	go f(ch2, ch3, "b")
	go f(ch3, ch1, "c")

	// 发送初始信号启动循环
	ch1 <- struct{}{}

	wg.Wait()
	close(ch1)
	close(ch2)
	close(ch3)
}
