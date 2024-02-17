package basic

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		if i <= 20 {
			fmt.Printf("======new goroutine======: i = %d\n", i)
			time.Sleep(1 * time.Second) //延时1s
		}
	}
}

func GoroutineDemo1() {
	//创建一个 goroutine，启动另外一个任务
	go newTask()
	i := 0
	//main goroutine 循环打印
	for {
		i++
		if i <= 20 {
			fmt.Printf("main goroutine: i = %d\n", i)
			time.Sleep(1 * time.Second) //延时1s
		}
	}
}

func GoroutineDemo2() {
	//创建一个 goroutine，启动另外一个任务
	go newTask()
	time.Sleep(1 * time.Second) //延时1s
	fmt.Println("main goroutine exit")
}
