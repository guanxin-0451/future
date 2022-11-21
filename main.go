package main

import (
	"fmt"
	"runtime"
)

//	[
//	{"id": 1, "name": "部门1", "pid": 0},
//	{"id": 2, "name": "部门2", "pid": 1},
//	{"id": 3, "name": "部门3", "pid": 1},
//	{"id": 4, "name": "部门4", "pid": 3},
//	{"id": 5, "name": "部门5", "pid": 4}
//]

type Partion struct {
	id   int
	name string
	pid  int
}

func initTree(partions []Partion) {

}

func main() {
	m := []int{1, 2, 3, 4, 5, 6, 7, 8}
	m1 := m[2:4]
	m = nil
	runtime.GC()
	fmt.Printf("%v, %v\n", m, m1[0:5])
	//ch := make(chan int, 10) //创建一个无缓存channel
	//
	////新建一个goroutine
	//
	//for i := 0; i < 10; i++ {
	//	ch <- i //往通道写数据
	//}
	////不需要再写数据时，关闭channel
	//close(ch)
	////ch <- 666 //关闭channel后无法再发送数据
	//
	//for {
	//	//如果ok为true，说明管道没有关闭
	//	if num, ok := <-ch; ok == true {
	//		fmt.Println("num = ", num)
	//	} else { //管道关闭
	//		fmt.Println("null")
	//	}
	//}
}
