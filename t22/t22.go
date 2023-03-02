/* Разработать программу, которая перемножает, делит, складывает, вычитает две
числовых переменных a,b, значение которых > 2^20.*/

package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	// при работе с числами > 2^20 потенциально можно выйти за границы типов int64
	// пакет big реализует арифметику произвольной точности (большие числа)

	// функция NewInt принимает значение int64
	b := big.NewInt(8223372036854775807)
	// есл int64 недостаточно для числа, можно получить значение из строки
	a, ok := new(big.Int).SetString("784563684183855136454135", 10)
	if !ok {
		log.Fatal("Error getting number from string")
	}
	fmt.Printf("a = %d\nb = %d\n", a, b)

	result := new(big.Int)
	fmt.Printf("Sum of numbers: %d\n", result.Add(a, b))
	fmt.Printf("Difference of numbers: %d\n", result.Sub(a, b))
	fmt.Printf("Multiplication of numbers: %d\n", result.Mul(a, b))
	fmt.Printf("Division of numbers: %d\n", result.Div(a, b))
}
