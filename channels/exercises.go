package channels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func couter(c chan int) {
	for i := 1; i <= 10; i++ {
		c <- i
		time.Sleep(1 * time.Second)
	}

	defer close(c)
}

func Exercise1() {
	c := make(chan int, 10)
	go couter(c)
	for i := range c {
		fmt.Println(i)
	}
}

type DownloadResult struct {
	ID      int
	Success bool
	Retries int
	Message string
}

func downloadImage(id int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	// mô phỏng thời gian download từ 1 đến 2 giây
	delay := time.Duration(rand.Intn(1000)+1000) * time.Millisecond
	time.Sleep(delay)

	ch <- fmt.Sprintf("📷 Ảnh #%d tải xong sau %.1f giây", id, delay.Seconds())
}

func Exercise2() {
	rand.Seed(time.Now().UnixNano())
	const total = 10

	ch := make(chan string, total)
	var wg sync.WaitGroup
	wg.Add(total)

	start := time.Now()

	for i := 1; i <= total; i++ {
		go downloadImage(i, ch, &wg)
	}

	// chờ tất cả goroutines hoàn thành
	wg.Wait()
	close(ch)

	fmt.Println("🎉 Kết quả:")
	for msg := range ch {
		fmt.Println(msg)
	}

	duration := time.Since(start).Seconds()
	fmt.Printf("\n⏱️ Tổng thời gian: %.2f giây\n", duration)
}
