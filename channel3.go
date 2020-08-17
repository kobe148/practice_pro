package main

import (
	"fmt"
	"math/rand"
	"time"
)

func source(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3)  + 1
	//睡眠1秒或者2秒或者3秒
	time.Sleep(time.Duration(rb) * time.Second)
	c <- ra
}

/**
最快回应：只要其中一个最快返回就采用这个结果，其他的就可以不用管。
 */
func main() {
	rand.Seed(time.Now().UnixNano())

	startTime := time.Now()
	c := make(chan int32, 5)//使用缓冲通道
	for i := 0; i < cap(c); i++ {
		go source(c)
	}
	rnd := <-c
	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)
}
