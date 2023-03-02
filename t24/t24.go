/* Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором.*/

package main

import (
	"fmt"
	"math"
)

// Point структура для хранения координат точки x и y
// координаты точки инкапсулируются за счет написания с маленькой буквы
// они не будут доступны из других пакетов
type Point struct {
	x float64
	y float64
}

// NewPoint конструктор структуры Point
func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

// GetPoint геттер для получения значений структуры Point
func (p Point) GetPoint() (float64, float64) {
	return p.x, p.y
}

func main() {
	point1 := NewPoint(3.5, -4.8)
	point2 := NewPoint(-6.3, 2.9)
	distance := getDistance(point1, point2)
	fmt.Printf("The distance between points is %.2f", distance)
}

// getDistance получает в качестве параметров 2 точки
// возвращает расстояние между ними
func getDistance(point1, point2 *Point) float64 {
	x1, y1 := point1.GetPoint()
	x2, y2 := point2.GetPoint()

	distance := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	return distance
}
