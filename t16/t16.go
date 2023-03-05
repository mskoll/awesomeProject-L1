/* Реализовать быструю сортировку массива (quicksort) встроенными методами
языка.*/

package main

import "fmt"

func main() {
	nums := []int{12, 6, 58, 96, 31, 24, 56, 87, 52, 91, 32, 64, 2, 3, 94, 23, 4, 10, 36, 78}
	fmt.Printf("%v\n", nums)
	quickSort(nums)
	fmt.Printf("%v\n", nums)
}

func quickSort(nums []int) []int {
	// базовый случай рекурсии
	if len(nums) < 2 {
		return nums
	}
	// границы сортировки, в то же время start высупает "разворотным",
	// end "опорным" индексами элементов
	start, end := 0, len(nums)-1

	// проходим по всем элементам сравнивая с опорным
	for i := start; i < end; i++ {
		// если текущий элемент меньше опорного
		if nums[i] < nums[end] {
			// меняем его местами с "разворотным" элементом
			nums[start], nums[i] = nums[i], nums[start]
			// смещаем разворотный элемент
			start++
		}
	}
	// меняем местами опорный и разворотный элементы
	nums[start], nums[end] = nums[end], nums[start]

	// сортируем левую часть, не включая опорный
	quickSort(nums[:start])
	// сортируем правую часть, не включая опорный
	quickSort(nums[start+1:])
	return nums
}
