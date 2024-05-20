package basic

import (
	"fmt"
	"time"
)

func GoroutineChannelDemo4_close_chan() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("In sub goroutine, 发送到--%d--到channel->c\n", i)
			c <- i
			//close(c)  	// panic: send on closed channel

		}
		fmt.Println("In sub goroutine, 关闭channel")
		close(c)
	}()

	time.Sleep(1 * time.Second)

	for {
		//ok为true说明channel没有关闭，为false说明管道已经关闭
		fmt.Println("===========接收channel->c的数据，并赋值给data")
		if data, ok := <-c; ok {
			time.Sleep(1 * time.Second)
			fmt.Printf("===========get data from channel->c: %d\n", data)
		} else {
			break
		}
	}

	fmt.Println("Finished")
}

func GoroutineChannelDemo4_range() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("In sub goroutine, 发送到--%d--到channel->c\n", i)
			c <- i
			//close(c)  	// panic: send on closed channel

		}
		fmt.Println("In sub goroutine, 关闭channel")
		close(c)
	}()

	//for {
	//	//ok为true说明channel没有关闭，为false说明管道已经关闭
	//	fmt.Println("===========接收channel->c的数据，并赋值给data")
	//	if data, ok := <-c; ok {
	//		fmt.Printf("===========get data from channel->c: %d\n", data)
	//	} else {
	//		break
	//	}
	//}
	for data := range c {
		time.Sleep(1 * time.Second)
		fmt.Printf("===========get data from channel->c: %d\n", data)
	}

	fmt.Println("Finished")
}
