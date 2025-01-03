package basic

import (
	"fmt"
	"time"
)

func GoroutineChannelDemo2_no_buffer() {

	var c chan int
	c = make(chan int, 0) //创建无缓冲的通道 c
	//c := make(chan int, 0) //创建无缓冲的通道 c

	//内置函数 len 返回未被读取的缓冲元素数量，cap 返回缓冲区大小
	fmt.Printf("len(c)=%d, cap(c)=%d\n", len(c), cap(c))

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 3; i++ {
			fmt.Printf("1, In sub goroutine, 发送--%d--到channel\n", i)
			c <- i
			fmt.Printf("4, 子go程正在运行[%d]: len(c)=%d, cap(c)=%d\n\n", i, len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second) //延时2s

	for i := 0; i < 3; i++ {
		fmt.Println("2, In main goroutine, 接收channel的数据，并赋值给num")
		num := <-c //从c中接收数据，并赋值给num
		fmt.Println("3, ===========num = ", num)
		//time.Sleep(1 * time.Second)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("===========main进程结束")
}
