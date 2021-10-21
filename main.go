package main

import (
	"fmt"
	"time"
)

var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

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

func exchange(nums []int) []int {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j]%2 == 0 {
			j++
		} else {
			swap(&nums, i, j)
			i++
		}
	}

	return nums
}

func swap(nums *[]int, i, j int) {
	n := *nums
	tmp := n[j]
	n[i] = n[j]
	n[j] = tmp
}
func main() {
	go write()
	go write2()
	read()
}
