package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	mu    sync.Mutex
	count int64
}

// 加锁
func (c *Counter) Increment(wg *sync.WaitGroup) {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			c.mu.Lock()
			defer c.mu.Unlock()
			for i := 0; i < 1000; i++ {
				c.count++
			}
			defer wg.Done()

		}()
	}
}

// 使用原子操作
func (c *Counter) IncrementAtomic(wg *sync.WaitGroup) {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {

			for i := 0; i < 1000; i++ {
				atomic.AddInt64(&c.count, 1)
			}
			defer wg.Done()

		}()
	}
}

func main() {
	var wg sync.WaitGroup
	c := Counter{
		mu:    sync.Mutex{},
		count: 0,
	}
	//c.Increment(&wg)
	c.IncrementAtomic(&wg)
	wg.Wait()
	fmt.Println(c.count)

}
