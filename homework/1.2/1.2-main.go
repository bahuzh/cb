package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i

		fmt.Println("producer ", i)
		time.Sleep(1 * time.Second)

	}
	close(ch)
}

func consumer(ch <-chan int) {
	for i := range ch {
		fmt.Println("consumer ", i)
	}
}

func main() {
	ch := make(chan int, 10) // var ch chan int,这样定义chan也可以，但是初始值时nil，不能使用，所以使用make关键字定义chan
	go producer(ch)          // go produce,让produce在协程中运行，这样不会阻塞consumer从chan中拿数据
	consumer(ch)

}
