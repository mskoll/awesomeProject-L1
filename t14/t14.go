/* Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.*/

package main

import (
	"fmt"
	"reflect"
)

func main() {

	i := 5
	s := "bazinga"
	b := true
	c := make(chan int)

	typeFmt(i)
	typeFmt(s)
	typeFmt(b)
	typeFmt(c)

	typeSwitch(i)
	typeSwitch(s)
	typeSwitch(b)
	typeSwitch(c)

	typeReflection(i)
	typeReflection(s)
	typeReflection(b)
	typeReflection(c)

}

func typeFmt(val interface{}) {
	// флаг %Т возращает представление типа согласно синтаксису Go
	fmt.Printf("<FMT> %T\n", val)
}

func typeSwitch(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Printf("<SWITCH> int\n")
	case string:
		fmt.Printf("<SWITCH> string\n")
	case bool:
		fmt.Printf("<SWITCH> bool\n")
	case chan int:
		fmt.Printf("<SWITCH> chan int\n")
	default:
		fmt.Printf("<SWITCH> unknown")
	}
}

func typeReflection(val any) {
	// помимо TypeOf можно использовать ValueOf().Kind()
	fmt.Printf("<REFLECT> %v\n", reflect.TypeOf(val))
}
