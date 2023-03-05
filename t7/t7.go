/* Реализовать конкурентную запись данных в map.*/

package main

import (
	"fmt"
	"sync"
)

// структура для мапы со встроенным мьютексом
type mutMap struct {
	sync.Mutex
	data map[int]int
}

func newMutMap() *mutMap {
	return &mutMap{data: make(map[int]int)}
}

// writeData функция обертка для записи данных в мапу
func (m *mutMap) writeData(key, value int) {
	m.Lock()
	m.data[key] = value
	m.Unlock()
}

func (m *mutMap) readAll() {
	for key, val := range m.data {
		fmt.Printf("%d - %d\n", key, val)
	}
}

func main() {
	mutMap := newMutMap()
	var wg sync.WaitGroup
	for i, j := 0, 100; i < 100; i, j = i+1, j-1 {
		wg.Add(1)
		go func(i, j int) {
			mutMap.writeData(i, j)
			wg.Done()
		}(i, j)
	}
	wg.Wait()
	fmt.Printf("Done\n")
	mutMap.readAll()
}
