package main

import (
	"fmt"
	"time"
)

var chs chan int = make(chan int, 2)
var chs2 chan int = make(chan int, 2)

func write() {
	//time.Sleep(3*time.Second)
	i := 1
	for {
		time.Sleep(time.Second)
		chs <- i
		i++
	}
}

func write2() {
	//time.Sleep(3*time.Second)
	i := 1
	for {
		time.Sleep(3 * time.Second)
		chs2 <- i
		i++
	}
}

func read() {
	for {
		select {
		case ch1 := <-chs:
			fmt.Printf("ch1: %d\n", ch1)
		case ch2 := <-chs2:
			fmt.Printf("ch2: %d\n", ch2)
			time.Sleep(10 * time.Second)
		}
	}

}

func main() {
	go write()
	go write2()
	read()
}
