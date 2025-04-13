package structs

import "fmt"

type Vertex struct {
	X, Y int
}

func BasicStruct() {
	v := Vertex{1, 2} // Create a struct instance
	p := &v // Create a pointer to the struct

	// v.X = 31 // Update age using value
	p.X = 31 // Update age using pointer

	fmt.Println(v)
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func LiteralStruct() {
	fmt.Println(v1, p, v2, v3)
}