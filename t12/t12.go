/*Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.*/

package main

import "fmt"

func main() {
	rows := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Printf("Set, created by slice: %v\n", createSetBySlice(rows))
	fmt.Printf("Set, created by map: %v\n", createSetByMap(rows))
}

// createSetBySlice возвращает срез уникальных строк,
// созданный путем удаления повторяющихся значений
func createSetBySlice(rows []string) []string {

	for i := 0; i < len(rows)-1; i++ {
		for j := i + 1; j < len(rows); j++ {
			// если в сресе есть повторное значение
			if rows[i] == rows[j] {
				// удаляем это значение из среза
				rows = append(rows[:j], rows[j+1:]...)
			}
		}
	}
	return rows
}

// createSetByMap возвращает срез уникальных строк,
// с помощью мапы
func createSetByMap(rows []string) (set []string) {
	rowsMap := make(map[string]struct{})
	// в качесте ключей мапы записываем все значения среза
	for _, row := range rows {
		rowsMap[row] = struct{}{}
	}
	// из полученной мапы записываем ключи в новый срез
	for key := range rowsMap {
		set = append(set, key)
	}
	return
}
