package generics

import (
	"fmt"
)

func mapAny[K, V any](arr []K, f func(K) V) []V {
	result := make([]V, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

func MapAny() {
	arr := []int{1, 2, 3, 4, 5}

	rs := mapAny(arr, func(i int) int {
		return i * 2
	})

	strArr := []string{"Apple", "Orange", "Lemon"}

	rsStr := mapAny(strArr, func(s string) string {
		return s + " fruit"
	})

	fmt.Println(rs)
	fmt.Println(rsStr)
}
