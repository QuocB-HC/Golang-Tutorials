package pointer

import "fmt"

func BasicPointer() {
	var p *int
	i := 42
	p = &i // p now points to i
	fmt.Println(*p) // Dereference p to get the value of i
	fmt.Println(p)  // Print the address of i
}

func AnotherBasicPointer() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}