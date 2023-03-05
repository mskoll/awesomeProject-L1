/* Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.*/

package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
	sync.Mutex
}

func newCounter() *counter {
	return &counter{count: 0}
}

func (c *counter) incr() {
	c.Lock()
	c.count++
	c.Unlock()
}

func (c *counter) getCount() int {
	return c.count
}

func main() {
	counter := newCounter()
	var wg sync.WaitGroup

	n := 1000
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			counter.incr()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Count: %d\n", counter.getCount())
}
