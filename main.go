package main

import (
	"fmt"
	// "test-app/interfaces"
	// "test-app/basic"
	"test-app/maps"
)

func main() {
	fmt.Println("Hello World!")

	// hypot := func(x, y float64) float64 {
	// 	return basic.SqrtNumber(x*x + y*y)
	// }

	// fmt.Println(hypot(5, 12))

	// fmt.Println(basic.FunctionValue(hypot))

	// add := basic.FunctionClosure()

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(
	// 		basic.FunctionClosure()(i), // Mỗi lần gọi hàm closure sẽ tạo ra một biến sum mới nên luôn luôn là 0 + i
	// 		add(i), // Đã tạo ra biến sum bên ngoài hàm closure nên mỗi lần gọi hàm closure sẽ cộng dồn vào biến sum
	// 	)
	// }

	// interfaces.TestInterface()
	// interfaces.EmptyInterface(5)
	// interfaces.TypeAssertion(true)
	// interfaces.Stringer()

	maps.MapInt()
}