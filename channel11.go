package main

import (
	"fmt"
	"time"
)

func Tick(d time.Duration) <-chan struct{} {
	c := make(chan struct{}, 1)
	go func() {
		for {
			time.Sleep(d)
			select {
			case c <- struct{}{}:
			default:
			}
		}
	}()
	return c
}

func main() {
	t := time.Now()
	//取出对应的数据，其实就是上面业务sleep一秒后
	for range Tick(time.Second) {
		fmt.Print(time.Since(t).Milliseconds())
	}
}
