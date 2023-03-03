/* Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть
регистронезависимой.
Например:
abcd — true
abCdefAaf — false
aabcd — false*/

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"
	fmt.Printf("%s - %v\n", str1, isUniqueBySlice(str1))
	fmt.Printf("%s - %v\n", str2, isUniqueBySlice(str2))
	fmt.Printf("%s - %v\n", str3, isUniqueBySlice(str3))

	fmt.Printf("%s - %v\n", str1, isUniqueByMap(str1))
	fmt.Printf("%s - %v\n", str2, isUniqueByMap(str2))
	fmt.Printf("%s - %v\n", str3, isUniqueByMap(str3))
}

func isUniqueBySlice(str string) bool {
	// приводим все символы к нижнему регистру
	str = strings.ToLower(str)
	// преобразуем строку к срезу рун
	runes := []rune(str)
	// перебираем все элементы среза
	for i := 0; i < len(runes)-1; i++ {
		for j := i + 1; j < len(runes); j++ {
			// если встречаем равные руны - возвращаем false
			if runes[i] == runes[j] {
				return false
			}
		}
	}
	return true
}

func isUniqueByMap(str string) bool {
	str = strings.ToLower(str)
	// инициализируем мапу
	strMap := make(map[rune]struct{})
	// перебираем все руны строки и записываем в мапу
	for _, r := range str {
		strMap[r] = struct{}{}
	}
	// если длина мапы и строки одинаковы - возвращаем true
	if len(strMap) == utf8.RuneCountInString(str) {
		return true
	}
	return false
}
