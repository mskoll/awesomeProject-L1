/* Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	fmt.Printf("%s - %s\n", str, reverse1(str))
	fmt.Printf("%s - %s\n", str, reverse2(str))
}

func reverse1(str string) string {
	// разбиваем строку на срез слов
	strs := strings.Split(str, " ")
	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		// симметрично смещаемся к центру, меняя элементы местами
		strs[i], strs[j] = strs[j], strs[i]
	}
	// преобразуем срез к строке
	return strings.Join(strs, " ")
}

func reverse2(str string) string {
	// разбиваем строку на срез слов
	strs := strings.Split(str, " ")
	var strB strings.Builder
	// записываем в билдер элементы, начиная с конца
	for i := len(strs) - 1; i >= 0; i-- {
		strB.WriteString(strs[i])
		strB.WriteString(" ")
	}
	// возвращаем строку
	return strB.String()
}
