package methods

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v Vertex) BasicMethod() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func MethodIsFunction(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) NonStructMethod() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Dùng con trỏ để thay đổi giá trị của struct
func (v *Vertex) MethodWithPointer(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Method
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Function
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func TestMethodAndPointerIndirection() {
	var v Vertex
	p := &v

	// Function
	//	ScaleFunc(v, 5)  // Compile error!
	ScaleFunc(&v, 5) // OK

	// Method
	v.Scale(5)  // OK
	p.Scale(10) // OK

	// Function
	fmt.Println(AbsFunc(v)) // OK
	//	fmt.Println(AbsFunc(&v)) // Compile error!

	// Method
	fmt.Println(v.Abs()) // OK
	fmt.Println(p.Abs()) // OK
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
