package main

import (
	"fmt"
	"time"
)

func fibonacii(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			//如果c可写，则该case就会进来
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	//sub go
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
			time.Sleep(1 * time.Second)
		}

		quit <- 0
	}()

	//main go
	fibonacii(c, quit)

	fmt.Println("end")
}
