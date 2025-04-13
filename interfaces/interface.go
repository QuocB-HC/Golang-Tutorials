package interfaces

import (
	"fmt"
	"math"
)

// Interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Struct hình chữ nhật
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Struct hình tròn
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func TestInterface() {
	var s Shape

	s = Rectangle{Width: 5, Height: 10}
	fmt.Println("Rectangle Area:", s.Area())
	fmt.Println("Rectangle Perimeter:", s.Perimeter())

	s = Circle{Radius: 7}
	fmt.Println("Circle Area:", s.Area())
	fmt.Println("Circle Perimeter:", s.Perimeter())
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func EmptyInterface(value any) {
	do(value)
}

func TypeAssertion(value any) {
	// Type assertion
	if str, ok := value.(string); ok {
		fmt.Println("String value:", str)
	} else if num, ok := value.(int); ok {
		fmt.Println("Integer value:", num)
	} else if float, ok := value.(float64); ok {
		fmt.Println("Float value:", float)
	} else if bool, ok := value.(bool); ok {
		fmt.Println("Boolean value:", bool)
	} else {
		fmt.Println("Unknown type")
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("\n%v (%v years)", p.Name, p.Age)
}

func Stringer() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
