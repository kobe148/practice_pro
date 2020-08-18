package main

import "fmt"

/**
通过用作互斥锁
 */

func main() {
	//如果是无缓冲channel会导致方法阻塞无法被别的协程取出数据
	mutex := make(chan struct{}, 1) //容量必须为1

	counter := 0
	increase := func() {
		mutex <- struct{}{} //加锁, 数据写入, 如果容量为0无缓冲的，这行会一直阻塞，不是产生死锁就是被别的协程读取数据
		counter++ //因为上面是有缓冲的，所以不会阻塞不执行，这行可以执行到
		<- mutex //解锁， 数据输出，channel空出可以继续写入数据
	}
	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			increase()
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done; <-done
	fmt.Println(counter) //2000
}
