package basic

import (
	"fmt"
)

// chan<- //只写
func counter(out chan<- int) {
	defer close(out)
	for i := 0; i < 10; i++ {
		fmt.Printf("1, In counter （chan<- int）, send only, 发送>>%d>>到channel->c\n", i)
		out <- i //如果对方不读 会阻塞
	}
}

// <-chan //只读
func printer(in <-chan int) {
	for num := range in {
		fmt.Printf("2, In printer （<-chan int）, read only, 接收<<%d<<自channel->c\n", num)
	}
}

func GoroutineChannelDemo5_one_direction() {
	//c := make(chan int)    //   chan   //读写
	//c := make(chan int, 1) //   chan   //读写
	//c := make(chan int, 2) //   chan   //读写
	c := make(chan int, 3) //   chan   //读写

	/*
		默认情况下，通道channel是双向的，也就是，既可以往里面发送数据也可以同里面接收数据。
		但是，我们经常见一个通道作为参数进行传递而只希望对方是单向使用的，要么只让它发送数据，要么只让它接收数据，这时候我们可以指定通道的方向。
		var ch1 chan int       // ch1是一个正常的channel，是双向的
		var ch2 chan<- float64 // ch2是单向channel，只用于写float64数据
		var ch3 <-chan int     // ch3是单向channel，只用于读int数据

		chan<- 表示数据进入管道，要把数据写进管道，对于调用者就是输出。
		<-chan 表示数据从管道出来，对于调用者就是得到管道的数据，当然就是输入。

		可以将 channel 隐式转换为单向队列，只收或只发，不能将单向 channel 转换为普通 channel：
		c := make(chan int, 3)
		var send chan<- int = c // send-only
		var recv <-chan int = c // receive-only

		send <- 1
		//<-send //invalid operation: <-send (receive from send-only type chan<- int)
		<-recv
		//recv <- 2 //invalid operation: recv <- 2 (send to receive-only type <-chan int)

	*/

	go counter(c) //生产者
	printer(c)    //消费者

	fmt.Println("done")
}
