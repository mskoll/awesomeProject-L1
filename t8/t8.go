/* Дана переменная int64. Разработать программу которая устанавливает i-й бит в
1 или 0.*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// исходное число
	var num int64 = 242 // 11110010
	i := 2
	// битовая маска, в нужном разряде
	var mask int64 = 1 << i // 100
	// используем исключающее или чтобы заменить нужный бит
	num = num ^ mask // 11110110
	fmt.Println(strconv.FormatInt(num, 2))
}
