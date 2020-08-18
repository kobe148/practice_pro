package main

import (
	"fmt"
	"time"
)

func main() {
	// 此信号通道也可以缓冲为1。如果这样，则在下面
	// 这个协程创建之前，我们必须向其中写入一个值。
	done := make(chan struct{})
	go func() {
		fmt.Print("hello")
		time.Sleep(time.Second * 2)

		<- done
	}()

	done <- struct{}{}
	fmt.Println(" world!")
}