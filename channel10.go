package main

import "runtime"

func DoSomething() {
	for {
		runtime.Gosched() //防止本携程霸占CPU不放
	}
}

func main() {
	go DoSomething()
	go DoSomething()
	select{}
}