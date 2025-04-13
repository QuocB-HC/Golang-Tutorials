package maps

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func MapExample() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

func MapExample2() {
	var m = map[string]Vertex{
		"Bell Labs": {
			40.68433,-74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}	

	fmt.Println(m)
}

func MulatingMap() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func mapInt(arr []int , f func(int) int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

func MapInt() {
	arr := []int{1, 2, 3, 4, 5}
	rs := mapInt(arr, func(i int) int {
		return i * 2
	})
	fmt.Println(rs)
}