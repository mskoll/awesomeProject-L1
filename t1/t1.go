/* Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования).*/

package main

import "fmt"

// Human структура с полями Name - имя, Surname - фамилия, Age - возраст
type Human struct {
	Name    string
	Surname string
	Age     uint
}

// NewHuman инициализатор(конструктор) структуры Human
func NewHuman(name, surname string, age uint) *Human {
	return &Human{Name: name, Surname: surname, Age: age}
}

// getFullName метод, возвращающий полное имя Human
func (h Human) getFullName() string {
	return fmt.Sprintf("%s %s", h.Name, h.Surname)
}

// getHumanInfo метод, возвращающий информацию о Human
func (h Human) getHumanInfo() string {
	return fmt.Sprintf("%s %s is %d years old", h.Name, h.Surname, h.Age)
}

// Action структура с полями Name и Human
// поле Human эквивалентно полю human Human
// но оно позволяет сократить путь к полю вложенной структуры
// например, вместо action.human.getHumanInfo() - action.getHumanInfo()
type Action struct {
	Name string
	Human
}

// getFullName метод, возвращающий полное имя Human
// структура Action содержит поле Name, идентичное полю Human
// a.Name - поле Action, a.Surname - поле Human
func (a Action) getFullName() string {
	return fmt.Sprintf("%s %s", a.Name, a.Surname)
}

func main() {
	human := NewHuman("John", "Doe", 27)
	action := Action{Name: "Jane", Human: *human}

	// можно напрямую вызывать методы вложенной структуры
	fmt.Printf("%s\n", action.getHumanInfo())

	// если у структур есть одинаковые методы - вызывается метод структуры Action
	fmt.Printf("%s\n", action.getFullName()) // Jane Doe

	// чтобы вызвать метод стуркуры Human - нужно явно обратиться к методу
	// через вложенную структуру
	fmt.Printf("%s\n", action.Human.getFullName()) // John Doe

	// также можно напрямую обращаться к полям вложенной структуры
	fmt.Printf("The human's surname is %s\n", action.Surname)

	// но при наличии одинаковых полей - берется поле структуры Action
	fmt.Printf("The action's name is %s\n", action.Name) // Jane

	// чтобы получить поле структуры Human - нужно явно обратиться к полю
	fmt.Printf("The human's name is %s\n", action.Human.Name) // John
}
