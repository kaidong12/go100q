package basic

import (
	"fmt"
	"time"
)

func fibonacci2(c chan int, quit chan int) {
	x, y := 1, 1
	for {
		// 与switch语句相比，select有比较多的限制，
		// 其中最大的一条限制就是每个case语句里必须是一个IO操作

		fmt.Println("x", x)
		select {
		case c <- x:
			x, y = y, x+y
			fmt.Println("x, y", x, y)
			//time.Sleep(1 * time.Second) //延时1s
		case <-quit:
			fmt.Println("quit")
			return
		}
		/*
			在一个select语句中，Go语言会按顺序从头至尾评估每一个发送和接收的语句。
			如果其中的任意一语句可以继续执行(即没有被阻塞)，那么就从那些可以执行的语句中任意选择一条来使用。
			如果没有任意一条语句可以执行(即所有的通道都被阻塞)，那么有两种可能的情况：
			- 如果给出了default语句，那么就会执行default语句，同时程序的执行会从select语句后的语句中恢复。
			- 如果没有default语句，那么select语句将被阻塞，直到至少有一个通信可以进行下去。
		*/

	}
}

/*
	select {
    case <- chan1:
        // 如果chan1成功读到数据，则进行该case处理语句
    case chan2 <- 1:
        // 如果成功向chan2写入数据，则进行该case处理语句
    default:
        // 如果上面都没有成功，则进入default处理流程
    }
*/

func GoroutineChannelDemo6_select_multiple_channel() {
	//c := make(chan int, 1)
	c := make(chan int)
	quit := make(chan int)

	go func() {
		//quit <- 0
		for i := 0; i < 8; i++ {
			fmt.Printf("In goroutine, read from c, %d -> %d\n", i, <-c)
		}
		time.Sleep(1 * time.Second) //延时1s
		quit <- 0
	}()

	fibonacci2(c, quit)

}
