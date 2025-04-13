package channels

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func BasicChannel() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum( s[ : len(s) / 2 ], c)
	go sum( s[ len(s) / 2 : ], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x + y)
}

func ChannelBuffer() {
	c := make(chan int, 2) // buffered channel with capacity 2
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}