package basic

import (
	"fmt"
)

func GoroutineChannelDemo1() {
	//var c chan int
	//c = make(chan int)

	c := make(chan int)

	go func() {
		defer fmt.Println("子go程结束")

		fmt.Println("子go程正在运行……")

		fmt.Println("In sub goroutine, 发送到666到channel: c <- 666")
		c <- 666 //666发送到c
	}()

	num := <-c //从c中接收数据，并赋值给num
	fmt.Println("In main goroutine, 接收channel的数据，并赋值给num: num := <-c")

	fmt.Println("===========num =", num)
	fmt.Println("===========main go程结束")
}
