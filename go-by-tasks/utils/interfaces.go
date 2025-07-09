package utils

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Println("Shape Area: ", s.Area())
}

// Empty Interface
func Describe(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}

func PrintType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Println("unknown type")
	}
}

// Interface Composition
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type Logger interface {
	Log(msg string)
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(msg string) {
	fmt.Println("LOG: ", msg)
}

func DoSomething(l Logger) {
	l.Log("Doing something....")
}

func InterfacesDemo() {
	shapes := []Shape{
		Rectangle{10, 4},
		Circle{5},
	}

	for _, s := range shapes {
		printArea(s)
	}
	Describe(shapes)
	PrintType(5)
	PrintType("test")

	log := ConsoleLogger{}
	DoSomething(log)
}
