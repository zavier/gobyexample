package main

import (
	"fmt"
	"time"
)

func worker(done chan bool)  {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	// 阻塞等待接收异步任务完成发送消息
	<-done
}