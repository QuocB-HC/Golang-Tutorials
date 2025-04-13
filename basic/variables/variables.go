package basic

import "fmt"

var c, python, java bool

var a, b int = 1, 2

func Variables() {
	var n int
	fmt.Println(n, c, python, java)

	var c, python, java = true, false, "no!"
	fmt.Println(a, b, c, python, java)

	var i, j int = 1, 2
	k := 3
	fmt.Println(i, j, k, c, python, java)
}
