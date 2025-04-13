package selects

import (
	"fmt"
	"time"
)

func sendEven(evenChan chan int) {
	num := 2
	for {
		time.Sleep(4 * time.Second)
		evenChan <- num
		num += 2
	}
}

func sendOdd(oddChan chan int) {
	num := 1
	for {
		time.Sleep(7 * time.Second)
		oddChan <- num
		num += 2
	}
}

func EvenAndOdd() {
	evenChan := make(chan int)
	oddChan := make(chan int)

	go sendEven(evenChan)
	go sendOdd(oddChan)

	// In ra 10 số nhận được từ 2 channel
	for i := 0; i < 10; i++ {
		select {
		case even := <-evenChan:
			fmt.Printf("🔵 Nhận số chẵn: %d\n", even)
		case odd := <-oddChan:
			fmt.Printf("🟠 Nhận số lẻ: %d\n", odd)
		}
	}
}