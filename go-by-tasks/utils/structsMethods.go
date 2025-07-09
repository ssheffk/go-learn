package utils

import "fmt"

// struct = custom data type (like a class without methods)

type Rectangle struct {
	Width, Height float64
}

type Person struct {
	Name string
	Age  int
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (p *Person) Rename(newName string) {
	p.Name = newName
}

// Embedding Structs (Composition)
type Address struct {
	City string
}
type Employee struct {
	Name string
	Address
}

// Method on Non-Struct Types
type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
	return float64(c)*1.8 + 32
}

func StructsMethodsDemo() {
	rect := Rectangle{10, 5}
	fmt.Println("Rectangle Area", rect.Area())

	// Anonymous Structs
	user := struct {
		ID   int
		Name string
	}{ID: 1, Name: "Temp"}
	fmt.Println("Anonymous Structs: ", user)

	e := Employee{Name: "Tom", Address: Address{City: "NYC"}}
	fmt.Println("Embedding Structs: ", e)

	fmt.Println("Non-Struct: ", Celsius.ToFahrenheit(5))
}
