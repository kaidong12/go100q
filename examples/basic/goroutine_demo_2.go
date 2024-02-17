package basic

import (
	"fmt"
	"runtime"
	"time"
)

func GoroutineDemo3() {

	// 匿名函数
	go func() {
		defer fmt.Println("A.defer")

		// 匿名函数
		func() {
			defer fmt.Println("B.defer")

			//退出当前goroutine
			//return
			runtime.Goexit() // 终止当前 goroutine, import "runtime"
			fmt.Println("B") // 不会执行
		}() //执行，不要忘记()

		fmt.Println("A") // 不会执行
	}() //执行，不要忘记()

	//死循环，目的不让主goroutine结束
	for {
		time.Sleep(10 * time.Second) //延时10s
		fmt.Println("Main")
	}
}

func GoroutineDemo4() {

	// 匿名函数
	go func(a int, b int) bool {
		fmt.Println("a= ", a, "b= ", b)
		defer fmt.Println("A.defer")

		// 匿名函数
		func(c string) {
			fmt.Println(c)
			defer fmt.Println("B.defer")

			//退出当前goroutine
			//return
			runtime.Goexit() // 终止当前 goroutine, import "runtime"
			fmt.Println("B") // 不会执行
		}("in sub-goroutine") //执行，不要忘记()

		fmt.Println("A") // 不会执行
		return true
	}(100, 200) //执行，不要忘记()

	//死循环，目的不让主goroutine结束
	for {
		time.Sleep(10 * time.Second) //延时10s
		fmt.Println("Main")
	}
}
