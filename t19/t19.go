/* Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "главрыба"
	fmt.Printf("Reverse string: %s\n", reverse1(str))
	fmt.Printf("Reverse string: %s\n", reverse2(str))

}

func reverse1(str string) string {
	// разбиваем строку на срез строк(символов)
	strSl := strings.Split(str, "")
	l := len(strSl) - 1
	for i := 0; i < l/2; i++ {
		// симметрично смещаемся к центру, меняя элементы местами
		strSl[i], strSl[l-i] = strSl[l-i], strSl[i]
	}
	// преобразуем срез к строке
	return strings.Join(strSl, "")
}

func reverse2(str string) string {
	// преобразуем строку к срезу рун
	runes := []rune(str)
	var strB strings.Builder
	// записываем в билдер руны, начиная с конца
	for i := len(runes) - 1; i >= 0; i-- {
		strB.WriteRune(runes[i])
	}
	// возвращаем строку
	return strB.String()
}
