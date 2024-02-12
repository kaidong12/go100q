package basic

import "fmt"

type shape interface {
	area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height

}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius

}

func Interface_demo() {

	var s shape

	s = Rectangle{10, 15}
	fmt.Printf("长方形面积：%f\n", s.area())

	s = Circle{10}
	fmt.Printf("长方形面积：%f\n", s.area())

}
