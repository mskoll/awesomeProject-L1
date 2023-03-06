/* Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при
старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
способ завершения работы всех воркеров.*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	var nWorkers int
	// цикл для ввода количества воркеров
	for {
		fmt.Printf("Enter the number of workers\n")
		_, err := fmt.Scan(&nWorkers)
		if err != nil {
			fmt.Printf("Incorrect value\n")
		} else {
			break
		}
	}

	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(nWorkers)

	// запуск воркеров
	for i := 0; i < nWorkers; i++ {
		go func(i int) {
			worker(ctx, i, ch)
			wg.Done()
		}(i)
	}

	// запись случайного числа в канал, завершающаяся по сигналу контекста
	go func() {
		for {
			select {
			case ch <- rand.Int():
			case <-ctx.Done():
				return
			}
			time.Sleep(time.Second)
		}
	}()

	// канал для сигналов системы
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	// ожидание сигнала
	<-stop
	fmt.Printf("Signal received\n")
	// отмена контекста
	cancel()
	wg.Wait()

	// мне кажется более правильно, тк по зпдпнию запись в главном потоке,
	// но оно не работает (не заверщается)
	// это запись в канал
	//for {
	//	select {
	//	case ch <- rand.Int():
	//
	//	case <-stop:
	//		fmt.Printf("Signal received\n")
	//		cancel()
	//		wg.Wait()
	//	}
	//	time.Sleep(time.Second)
	//}
}

func worker(ctx context.Context, id int, ch1 chan int) {
	for {
		select {
		// вывод полученных данных
		case msg := <-ch1:
			fmt.Printf("Worker %d recived %d\n", id, msg)
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped\n", id)
			return
		}
	}
}
