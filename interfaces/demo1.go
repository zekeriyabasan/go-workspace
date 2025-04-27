package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangular struct {
	height float64
	width  float64
}

type circle struct {
	radius float64
}

func (r rectangular) area() float64 {
	return r.height * r.width
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func school(s shape) {
	fmt.Println("Alanı : ", s.area())
}

func Demo1() {
	r := rectangular{width: 10.323, height: 323.103}
	school(r)
	c := circle{radius: 21.3205}
	school(c)

}
