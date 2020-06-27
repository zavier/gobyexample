package main

import "fmt"

func main() {
	messages := make(chan string)

	// 不异步执行会死锁
	go func() {messages <- "ping"}()

	msg := <-messages
	fmt.Println(msg)
}