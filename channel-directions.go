package main

import "fmt"

// 声明pings channel只能用来发送消息，不能接收
func ping(pings chan<- string, msg string) {
	pings <- msg
}


func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
