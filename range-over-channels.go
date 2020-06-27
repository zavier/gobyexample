package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	// 不关闭的话下面的range会阻塞死锁
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}