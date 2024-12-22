package basic

import (
	"fmt"
	"sync"
	"time"
)

var group sync.WaitGroup

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("hello......")
		time.Sleep(time.Second)
	}
	//线程结束 -1
	group.Done()
}

func sayHi() {
	//线程结束 -1
	defer group.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("hi......")
		time.Sleep(time.Second)
	}
}

func wait_demo() {
	//+2
	group.Add(2)
	fmt.Println("main正在阻塞...")
	go sayHello()
	fmt.Println("main持续阻塞...")
	go sayHi()
	//线程等待
	group.Wait()
	fmt.Println("main貌似结束了阻塞...")
}
