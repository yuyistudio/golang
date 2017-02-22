package main

import (
	"fmt"
	"time"
)

func p1(c chan int) {
	for i := 0; i < 10; i++ {
		select {
		case c <- i:
			fmt.Println("send %v", i)
		default:
			fmt.Println("pass %v", i)
		}
	}
}

func p2(c chan int) {
	for {
		time.Sleep(200 * time.Millisecond)
		v := <- c
		fmt.Println("recv %v", v)
	}
}

func main() {
	q := make(chan int, 5)
	go p2(q)
	go p1(q)
	time.Sleep(2 * time.Second)
	fmt.Println(len(q))
}

