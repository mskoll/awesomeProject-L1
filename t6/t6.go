/* Реализовать все возможные способы остановки выполнения горутины.*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	chanClose()
	ctxCancel()
}

// chanClose остановка горутины через канал
func chanClose() {
	ch := make(chan int)
	stop := make(chan struct{})
	go func() {
		fmt.Printf("Goroutine started\n")
		for {
			select {
			// вывод данных, полученных из ch
			case c := <-ch:
				fmt.Printf("%d\n", c)
			case <-stop:
				// закрытие ch
				close(ch)
				// выход из горутины
				return
			}
		}
	}()
	for i := 0; i < 10; i++ {
		// запись данных в канал
		ch <- i
		time.Sleep(time.Millisecond * 300)
	}
	// отправка данных для остановки горутины
	stop <- struct{}{}
	fmt.Printf("Goroutine stopped\n")
}

// ctxCancel остановка горутины с помощью контекста
func ctxCancel() {
	// канал для сигнала остановки горутины
	ch := make(chan struct{})
	// WithCancel/WithDeadline/WithTimeout
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		fmt.Printf("Goroutine started\n")
		for {
			select {
			case <-ctx.Done():
				// отправка данных об остановке горутины
				ch <- struct{}{}
				return
			}
		}
	}(ctx)
	go func() {
		// завершаем горутины через 3 секунды
		time.Sleep(time.Second * 3)
		cancel()
	}()
	// получение даннных об остановке горутины
	<-ch
	fmt.Printf("Goroutine stopped\n")
}
