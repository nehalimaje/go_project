package main

import (
	"fmt"
	"math"
	"reflect"
)

type Shape interface {
	Area() float32
	Peri() float32
	Temp() string
}
type Circle struct {
	Red float32
}

type Square struct {
	L float32
	B float32
}

func (c *Circle) Temp() string {
	return ""
}
func (c *Square) Temp() string {
	return ""
}
func (s *Square) Area() float32 {
	fmt.Println("Square Area method is called ")
	return s.L * s.B
}

func (s *Square) Peri() float32 {
	fmt.Println("Square Peri method is called ")
	return s.L * s.B
}

func (c *Circle) Area() float32 {
	fmt.Println("Circle Area method is called ")
	return c.Red * math.Pi * c.Red
}
func (c *Circle) Peri() float32 {
	fmt.Println("Circle Peri method is called ")
	return c.Red * math.Pi * c.Red
}
func Calculate(a Shape) {
	fmt.Println("TYPE OF VARIABLE IS ", reflect.TypeOf(a))
	a.Area()
}
func main() {
	fmt.Println("Start")
	sq := Square{L: 1.2, B: 1.2}
	fmt.Println("Squre Area is  ", sq.Area(), sq.Peri())
	c := Circle{Red: 1.2}
	fmt.Println("Circle Area is  ", c.Area())
	Calculate(&sq)
	Calculate(&c)
}

