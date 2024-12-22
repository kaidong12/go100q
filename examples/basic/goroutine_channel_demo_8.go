package basic

import (
	"fmt"
	"math/rand"
	"time"
)

func GoroutineChannelDemo8_select_timeout() {
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go sleepyGooher(i, c)
	}

	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("gopher", gopherID)
		case <-timeout:
			fmt.Println("My patience run out")
			return
		}
	}
}

func sleepyGooher(id int, c chan int) {
	sleep_time := rand.Intn(4000)
	fmt.Printf("sleep time for sleepyGooher %d is %d\n", id, sleep_time)
	time.Sleep(time.Duration(sleep_time) * time.Millisecond)
	c <- id
}
