package flowcontrols

import "fmt"

func BasicFor()  {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Print(sum)
		fmt.Print(" ")
	}
	fmt.Println()

	sum = 1
	for ; sum < 1000 ; {
		sum += sum
		fmt.Print(sum)
		fmt.Print(" ")
	}
	fmt.Println()
}

func WhileLoop() {
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Print(sum)
		fmt.Print(" ")
	}
}

func InfiniteLoop() {
	sum := 0
	for {
		sum ++
		fmt.Print(sum)
		fmt.Print(" ")

		if sum > 1000 {
			break
		}
	}
}
