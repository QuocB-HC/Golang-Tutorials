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

func ForWithRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}

	fmt.Println("\nUsing _ to ignore index")
	for _, v := range pow {
		fmt.Println(v)
	}

	fmt.Println("\nUsing _ to ignore value")
	for i := range pow {
		fmt.Println(i)
	}
}