package main

import (
	"fmt"
	"math"
	"reflect"
)

type Shape interface {
	Area() float32
	GetName() string
	Peri() float32
}
type Square struct {
	L float32
	B float64
}

type Circle struct {
	Red float32
}

func (crl *Circle) Area() float32 {

	return math.Pi * 2 * crl.Red * crl.Red
}

func (ccc *Circle) GetName() string {
	return "NAME : Circle"
}

func (ca *Circle) Peri() float32 {
	return ca.Red
}
func (s *Square) GetName() string {
	return "SQUARE IS NAME"
}

func (s *Square) Peri() float32 {
	return 1.1
}

func (s *Square) Area() float32 {

	return float32(s.B) * s.L
}
func TestVar(a int, list ...int) {

}

func DynamicAreaCalling(s Shape) {
	fmt.Println("")
	fmt.Println("Type", reflect.TypeOf(s))
	fmt.Println("Dynami calling is ", s.Area())
}
func main() {
	fmt.Println("Start")
	ccVariable := &Circle{Red: 1.13}
	fmt.Println("ccVariable.Area()", ccVariable.Area())
	fmt.Println("GetName", ccVariable.GetName())
	fmt.Println(" Peri ", ccVariable.Peri())
	// ccVariable.TestData()

	////////////////
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	sqr := &Square{L: 4.6, B: 5.7}
	// fmt.Println("Area", sqr.Area())
	fmt.Println("GetName", sqr.GetName())
	// fmt.Println("Peri", sqr.Peri())

	////////////
	fmt.Println("")
	DynamicAreaCalling(sqr)
	DynamicAreaCalling(ccVariable)

}

func (c *Circle) TestData() float32 {
	return 4.5
}
