package goroutines

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Second)
	}
}

func SaySomething() {
	go say("Hello")
	say("World")
}
