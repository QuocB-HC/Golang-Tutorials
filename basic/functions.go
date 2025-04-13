package basic


func AddFunction(x, y int) int {
	return x + y
}

func SwapFunction(x, y string) (string, string) {
	return y, x
}

func SplitFunction(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
