package main

import (
	"fmt"
	"sync"
)

// golang交替打印奇偶数
func printOddEven(num int) {
	var ch = make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= num; i += 2 {
			<-ch
			if i%2 == 1 {
				fmt.Println("odd:", i)
			}
			ch <- struct{}{}
		}
		<-ch
	}()
	go func() {
		defer wg.Done()
		ch <- struct{}{}
		for i := 2; i <= num; i += 2 {
			<-ch
			if i%2 == 0 {
				fmt.Println("even:", i)
			}
			ch <- struct{}{}
		}
	}()

	wg.Wait()
	fmt.Println("end")
	close(ch)
}
