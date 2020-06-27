package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i < 5; i++ {
		// 类似Java中的CountdownLatch, Add的数量和调用Done的次数要一致
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

}