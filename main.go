package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)

	go ping(ch1)
	go pong(ch1)

	time.Sleep(time.Second * 10)
	x := <-ch1
	fmt.Println(x)
}

func ping(ch chan int) {
	i := 1
	for {
		fmt.Println("PING!")
		time.Sleep(time.Second / 2)
		ch <- i + 1
		i = <-ch
	}
}

func pong(ch chan int) {
	for {
		ii := <-ch
		fmt.Println("PONG!")
		time.Sleep(time.Second / 2)
		ch <- ii + 1
	}
}
