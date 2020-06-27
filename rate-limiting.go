package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 等于 limiter := time.NewTicker(d).C
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		// 200毫秒通过一次
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 限流器
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			// 定时200毫秒一次
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		// 限流器使用
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}