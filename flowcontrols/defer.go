package flowcontrols

import "fmt"

func BasicDefer() {
	fmt.Println("defer example")

	defer fmt.Println("world")
	fmt.Print("hello ")
}

func CountingDefer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
