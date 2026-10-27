package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	rectangle := Rectangle{Width: 10, Height: 20}
	circle := Circle{Radius: 2}

	fmt.Printf("长方形的周长: %f, 面积: %f \n", rectangle.Perimeter(), rectangle.Area())
	fmt.Printf("圆形的周长: %f, 面积: %f \n", circle.Perimeter(), circle.Area())
}
