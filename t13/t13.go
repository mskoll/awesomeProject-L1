/* Поменять местами два числа без создания временной переменной.*/

package main

import "fmt"

func main() {

	num1 := 4
	num2 := 9
	fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)

	num1, num2 = num2, num1
	fmt.Printf("First swap:\tnum1 = %d, num2 = %d\n", num1, num2)

	num1 = num1 + num2 // num1 = 4 + 9 	= 13
	num2 = num1 - num2 // num2 = 13 - 9 = 4
	num1 = num1 - num2 // num3 = 13 - 4 = 9
	fmt.Printf("Second swap:\tnum1 = %d, num2 = %d\n", num1, num2)

	num1 = num1 ^ num2 // num1 = 0100 ^ 1001 = 1101
	num2 = num1 ^ num2 // num2 = 1101 ^ 1001 = 0100
	num1 = num1 ^ num2 // num1 = 1101 ^ 0100 = 1001
	fmt.Printf("Third swap:\tnum1 = %d, num2 = %d\n", num1, num2)
}
