package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest() <-chan int32 {
	//无缓冲通道
	r := make(chan int32)

	go func() {
		time.Sleep(time.Second * 3)
		//写入一个随机数
		r <- rand.Int31n(100)
	}()

	return r
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//每个通道读取操作将阻塞到请求返回结果为止
	//a, b 相当于 <-50 <-30的形式被写入channel，然后会被阻塞，需要将其写出才不会阻塞
	a, b := longTimeRequest(), longTimeRequest()
	fmt.Println(sumSquares(<-a, <-b))
}
