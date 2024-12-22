package basic

import (
	"fmt"
	"time"
)

func GoroutineChannelDemo4_close_chan_read() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("1, In sub goroutine, 发送--%d--到channel->c\n", i)
			c <- i
			//close(c)  	// panic: send on closed channel

		}
		fmt.Println("4, In sub goroutine, 关闭channel")
		close(c)
	}()

	time.Sleep(1 * time.Second)

	for {
		//ok为true说明channel没有关闭，为false说明管道已经关闭
		fmt.Println("2, ===========接收channel->c的数据，并赋值给data")
		if data, ok := <-c; ok {
			time.Sleep(1 * time.Second)
			fmt.Printf("3, ===========get data from channel->c: %d\n\n", data)
		} else {
			fmt.Println("4, Receiving Closed Channel:")
			fmt.Println("4, When you receive from a closed channel, the value will be the zero value of the channel's element type. ")
			fmt.Println("4, For example, if the channel is of type `chan int`, the zero value is `0`. If the channel is of type `chan string`, the zero value is an empty string `\"\"`. ")
			fmt.Println("4, Additionally, the `ok` return value will be `false`, indicating that the channel is closed.")
			fmt.Printf("\n4, data=%d, ok=%t \n", data, ok)
			break
		}
	}

	fmt.Println("Finished")
}

func GoroutineChannelDemo4_close_chan_range() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for num := range ch {
		fmt.Println(num)
	}
}
