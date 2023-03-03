/* Реализовать собственную функцию sleep.*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	sleep(5)
	sleepCtx(2)
}

func sleep(duration time.Duration) {
	fmt.Printf("Start sleeping %d seconds\n", duration)

	// time.After возвращает канал, по истечении указанной продолжительности
	<-time.After(time.Second * duration)

	fmt.Printf("Finish sleeping\n")
}

func sleepCtx(duration time.Duration) {
	fmt.Printf("Start sleeping %d seconds\n", duration)

	// создаем контекст с заданным временем жизни
	ctx, _ := context.WithTimeout(context.Background(), time.Second*duration)
	<-ctx.Done() // получаем сигнал, когда контекст истек

	fmt.Printf("Finish sleeping\n")
}
