/* Реализовать пересечение двух неупорядоченных множеств.*/

package main

import "fmt"

func main() {

	set1 := []int{25, 45, 63, 12, 48, 94, 71, 35, 26, 14}
	set2 := []int{78, 53, 21, 65, 95, 48, 25, 12, 94, 63, 15, 26}
	fmt.Printf("Set intersection result by slice: %v\n", intersectionSlice(set1, set2))
	fmt.Printf("Set intersection result by map: %v\n", intersectionMap(set1, set2))

}

// intersectionSlice пересечение множеств путем пееребора
func intersectionSlice(set1, set2 []int) (set []int) {
	for _, el1 := range set1 {
		for _, el2 := range set2 {
			// если элемент из первого множества присутствует
			// во втором множестве
			if el1 == el2 {
				// записываем его в итоговый срез
				set = append(set, el1)
				break
			}
		}
	}
	return
}

// intersectionMap пересечение множеств с помощью мапы
func intersectionMap(set1 []int, set2 []int) (set []int) {
	// мапа для хранения значений множеств
	// ключ - значение из множества
	// значение - кочисество значений
	setMap := make(map[int]int)
	for _, el := range set1 {
		setMap[el]++
	}
	for _, el := range set2 {
		setMap[el]++
	}
	for key, count := range setMap {
		// если значение из множества встретилось более 1 раза
		if count > 1 {
			// добавляем его в итоговый срез
			set = append(set, key)
		}
	}
	return
}
