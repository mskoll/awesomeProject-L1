/* Удалить i-ый элемент из слайса.*/

package main

import "fmt"

func main() {

	i := 4
	fmt.Printf("Result: %v\n", deleteEl1(i))
	fmt.Printf("Result: %v\n", deleteEl2(i))

}

// deleteEl1 удаление элемента с сохранением порядка
// дорогостоящий вариант, т.к. приходится перемещать элементы
func deleteEl1(i int) []int {
	sl := []int{5, 4, 8, 6, 9, 2, 3, 8}
	return append(sl[:i], sl[i+1:]...)
}

// deleteEl2 удаление элемента без сохранения порядка
// быстрый вариант, выполняется за постоянное время
func deleteEl2(i int) []int {
	sl := []int{5, 4, 8, 6, 9, 2, 3, 8}
	sl[i] = sl[len(sl)-1]
	return sl[:len(sl)-1]
}
