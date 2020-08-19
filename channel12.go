package main

import (
	"fmt"
	"time"
)

/**
速率限制
 */
type Reuqest interface {}
func handle(r Reuqest) {fmt.Println(r.(int))}

const RateLimitPeriod = time.Minute
const RateLimit = 200

func handleRequest(requests <-chan Reuqest) {
	quotas := make(chan time.Time, RateLimit)

	go func() {
		tick := time.NewTicker(RateLimitPeriod / RateLimit)
		defer tick.Stop()
		for t := range tick.C {
			select {
			case quotas <- t:
			default:
			}
		}
	}()

	for r := range requests {
		<- quotas
		go handle(r)
	}
}

func main() {
	requests := make(chan Reuqest)
	go handleRequest(requests)
	for i := 0; ; i++ {requests <- i}
}
