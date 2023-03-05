/* Реализовать паттерн «адаптер» на любом примере.*/
package main

import "fmt"

// основной интерфейс, для вывода слонов
type printer interface {
	printElephant()
}

// сервис, который умеет печатать слонов
type elephant struct {
	eleph string
}

func (el *elephant) printElephant() {
	fmt.Printf("%s\n\n\n", el.eleph)
}

// адаптер, превращающий муху в слона
type adapter struct {
	fly *fly
}

func (a *adapter) printElephant() {
	a.fly.printFly()
}

// сервис, который умеет печатать мух, но не умеет печатать слонов
type fly struct {
	f string
}

func (f *fly) printFly() {
	fmt.Printf("%s\n", f.f)
}

// клиент, который хочет получить муху
type client struct{}

func (c *client) print(print printer) {
	print.printElephant()
}

func main() {
	elephPic := "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣯⡩⠀⠀\n⠀⠀⠀⡠⠔⠒⠒⠤⠤⢤⣴⠾⢿⠋⠑⢤⡱⣝⢤⠀\n⠀⡠⠊⠀⠀⠀⠀⠀⢠⣏⠃⠀⠘⠆⠰⠇⠑⠞⠀⡇\n⡜⠁⠀⠀⠀⠀⠀⠀⠀⢹⢦⡀⠀⠀⠀⢠⡤⣀⣴⣵\n⡇⠀⠀⠀⢠⡆⠀⠀⠀⡜⠀⠁⢀⣀⡗⠊⠙⠉⠉⠁\n⣿⠀⠀⠀⣸⡁⠀⣀⣠⡇⠀⠀⢸⠀⠀⠀⠀⠀⠀⠀\n⣿⠀⠀⣰⢫⠋⠉⠀⢸⠳⡄⠀⡎⠀⠀⠀⠀⠀⠀⠀\n⢹⠀⢰⡇⢸⠀⠀⠀⡇⠀⢳⠀⡇⠀⠀⠀⠀⠀⠀⠀\n⢸⣀⣸⠧⠬⠂⠀⠐⠥⡸⢸⣀⣳⠀⠀⠀⠀⠀⠀⠀"
	flyPic := "⠀⠀⠀⠀Я СЛОН⠀⠀⠀⠀⠀\n⠐⠲⡀⠀⠀⢐⣦⣂⠀⠀⠴⠀⠀⠀⠀\n⠀⠀⠨⣿⡈⣹⣂⣮⠉⢷⠁⠀⠀⠀⠀\n⠀⠀⠀⢻⡽⠭⠧⡧⠯⢱⠀⠀⠀⠀⠀\n⠀⠀⠀⠾⠀⡀⡁⠇⢤⣸⣂⣀⢀⣠⠀\n⠀⠀⡜⢐⢙⢯⣤⡤⣏⠀⡏⠁⠀⠀⠀\n⠀⠸⠪⣷⠈⠀⠸⡆⠀⠸⣷⡆⠀⠀⠀\n⠀⠶⣶⣿⠸⢠⣰⠂⠘⣰⠙⣷⠀⠀⠀\n⠀⣀⡟⢣⠑⣥⡃⢈⡴⠹⠀⠉⡛⠖⠀\n⠀⡅⠃⠄⢀⠚⠁⠊⠂⡀⠀⠈⠇⠀⠀\n⠀⠈⠪⠜⠁⠀⠀⠀⠀⠀⠢⠌⠀⠀⠀⠀"
	client := &client{}
	elephant := elephant{eleph: elephPic}
	client.print(&elephant)
	fly := fly{f: flyPic}
	adapter := &adapter{fly: &fly}
	client.print(adapter)
}
