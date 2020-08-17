package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequestTwo(r chan<- int32) {
	time.Sleep(time.Second * 3)
	r <-rand.Int31n(100)
}

func sumSquaresTwo(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ra, rb := make(chan int32), make(chan int32)
	go longTimeRequestTwo(ra)
	go longTimeRequestTwo(rb)

	//如果这里要打印出变量的值，不定义变量接收通道返回的值传入参数会死锁，因为通道已经关闭了。
	a := <-ra
	b := <-rb
	fmt.Println(a)
	fmt.Println(b)
	//如果不打印出通道的值可以直接使用通道传参数
	//fmt.Println(sumSquaresTwo(<-ra, <-rb))
	fmt.Println(sumSquaresTwo(a, b))
}