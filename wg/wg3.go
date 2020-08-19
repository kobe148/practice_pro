package main

import (
	"log"
	"sync"
)

func main() {
	log.SetFlags(0)

	x := 0
	doSometing := func() {
		x++
		log.Println("hello")
	}
	var wg sync.WaitGroup
	var once sync.Once
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(doSometing)
			log.Println("world!")
		}()
	}
	wg.Wait()
	log.Println("x = ", x)
}
