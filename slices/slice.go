package slices

import (
	"fmt"
	"strings"
)

func BasicSlice() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func SliceDefault() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func SliceLengthAndCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// A nil slice has a length and capacity of 0 and has no underlying array.
func NilSlice() {
	var s []int
	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil!")
	}
}

func SliceWithMake() {
	a := make([]int, 5)
	printSliceWithMake("a", a)

	b := make([]int, 0, 5)
	printSliceWithMake("b", b)

	c := b[:2]
	printSliceWithMake("c", c)

	d := c[2:5]
	printSliceWithMake("d", d)
}

func printSliceWithMake(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func JoinInString() {
	word := []string{"Go", "is", "awesome"}
	fmt.Println(word) // normal print

	result := strings.Join(word, " ")
	fmt.Println(result)                  // join with space
	fmt.Println(strings.Join(word, "-")) // join with dash
}

func SlicesOfSlices() {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	fmt.Println()
	fmt.Printf("%s\n", strings.Join(board[0], " "))
	fmt.Println(board[0][0]) // [] đầu lây dọc, [] sau lấy ngang
	fmt.Println(board[1][0]) // X
	fmt.Println(board[0][1]) // _
}

func AppendSlice() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}