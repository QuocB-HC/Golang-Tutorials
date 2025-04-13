package goroutines

import (
    "fmt"
    "sync"
    "time"
)

func says(name string, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 3; i++ {
        fmt.Println(name)
        time.Sleep(1 * time.Second)
    }
}

func SayWithWaitGroup() {
    var wg sync.WaitGroup
    wg.Add(2) // đợi 2 goroutines

    go says("Goroutine 1", &wg)
    go says("Goroutine 2", &wg)

    wg.Wait() // đợi đến khi cả 2 goroutine xong
    fmt.Println("All done!")
}