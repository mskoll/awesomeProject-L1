/* Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	var n time.Duration = 1
	ctx, _ := context.WithTimeout(context.Background(), time.Second*n)

	ch := make(chan int)
	go writer(ch)
	go reader(ch)
	// получаем сигнал, когда контекст истек
	<-ctx.Done()

}

// writer запись данных в канал
func writer(ch chan<- int) {
	i := 0
	for {
		ch <- i
		i++
	}
}

// reader чтение данных из канала
func reader(ch <-chan int) {
	for c := range ch {
		fmt.Printf("CHAN VALUE %d\n", c)
	}
}
