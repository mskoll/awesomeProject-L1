/* К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.
var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}
func main() {
	someFunc()
}*/

package main

import (
	"fmt"
	"strings"
)

var justString string

func main() {
	someFunc()
	otherFunc()
}

// someFunc работает некорректно, т.к. срез идет по количеству байт, а не рун
// 1 байт != 1 руне, ввиду чего мы можем потерять часть информации
// вывод: もしも�
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
	fmt.Printf("SOMEFUNC %s\n", justString)
}

// otherFunc делает срез по количеству рун, данные не теряются
func otherFunc() {
	v := createHugeString(1 << 10)
	justString = string([]rune(v)[:100])
	fmt.Printf("OTHERFUNC %s\n", justString)
}

func createHugeString(size int) string {
	var sBuild strings.Builder

	for i := 0; i < size; i++ {
		sBuild.WriteString("もしもし")
	}
	return sBuild.String()
}
