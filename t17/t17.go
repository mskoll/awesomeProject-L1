/* Реализовать бинарный поиск встроенными методами языка.*/

package main

import (
	"fmt"
	"sort"
)

func main() {

	numbers := []int{1, 3, 5, 6, 8, 9, 10, 12, 14, 15, 17, 20}
	target1 := -5 // i = 3
	target2 := 12 // i = 7
	target3 := 17 // i = 10
	target4 := 7  // i = -1

	fmt.Printf("Element %d is at index %d\n", target1, binarySearchRecursive(numbers, target1))
	fmt.Printf("Element %d is at index %d\n", target2, binarySearchRecursive(numbers, target2))

	fmt.Printf("Element %d is at index %d\n", target3, binarySearchIterative(numbers, target3))
	fmt.Printf("Element %d is at index %d\n", target4, binarySearchIterative(numbers, target4))

	// в Go есть встроенная функция бинарного поиска sort.Search()
	// но данная функция предполагает наличие искомого элемента в срезе
	// если искомый элемент может отсутствовать в срезе - его начилие нужно проверить отдельно
	// при отсутствии искомого элемента в срезе, функция возвращает индекс, куда он мог бы быль вставлен
	target5 := 13
	ind := sort.Search(len(numbers), func(i int) bool {
		// если срез отсортирован в обратном порядке, нужно использовать <=
		return numbers[i] >= target5
	})
	if ind < len(numbers) && numbers[ind] == target5 {
		fmt.Printf("Element %d is at index %d\n", target5, ind)
	} else {
		ind = -1
		fmt.Printf("Element %d is at index %d\n", target5, ind)
	}

}

// binarySearchRecursive рекурсивный бинарный поиск
// nums - срез чисел, target - искомое число
// функция возвращает ind - индекс элемента в срезе (-1 - элемент не найден)
func binarySearchRecursive(nums []int, target int) (ind int) {

	// если срез исчерпан - элемент не найден
	if len(nums) == 0 {
		return -1
	}

	// mid - середина среза
	mid := len(nums) / 2

	// если элемент найден - возвращаем индекс mid
	if target == nums[mid] {
		ind = mid
		// если элемент меньше среднего элемента среза
	} else if target < nums[mid] {
		// вызываем функцию для среза от 0 до mid
		ind = binarySearchRecursive(nums[:mid], target)
		// иначе (если элемент больше среднего элемента среза)
	} else {
		// вызываем функцию для среза от mid+1 до конца среза
		ind = binarySearchRecursive(nums[mid+1:], target)
		// если элемент был найден
		if ind > -1 {
			// прибавляем к текущему индексу значение среднего индекса из рекурсивного вызова
			ind += mid + 1
		}
	}
	return
}

// binarySearchRecursive рекурсивный бинарный поиск
// nums - срез чисел, target - искомое число
// функция возвращает ind - индекс элемента в срезе (-1 - элемент не найден)
func binarySearchIterative(nums []int, target int) (ind int) {
	firstInd := 0
	lastInd := len(nums) - 1

	// цикл, пока срез не закончится
	for firstInd <= lastInd {
		// определяем середину среза
		mid := (firstInd + lastInd) / 2
		// если элемент найден - возвращаем индекс элемента
		if target == nums[mid] {
			return mid
			// если элемент меньше среднего элемента среза
		} else if target < nums[mid] {
			// смещаем индекс правой границы среза
			lastInd = mid - 1
		} else {
			// смещаем индекс левой границы среза
			firstInd = mid + 1
		}
	}
	return -1
}
