package basic

import (
	"fmt"
	"time"
)

type CooldownService2 struct {
	cooldownDuration time.Duration
	timeout          <-chan time.Time
	cooldownCh       chan struct{}
	isCooldownActive bool
}

func NewCooldownService2(cooldownDuration time.Duration) *CooldownService {
	return &CooldownService{
		cooldownDuration: cooldownDuration,
		cooldownCh:       make(chan struct{}),
	}
}

func (cs *CooldownService2) Start2() {
	var ct time.Time
	if cs.isCooldownActive {
		fmt.Println("Start():", "1, Already in cooldown, please wait.")
		fmt.Printf("Start(): 2, Start time: %d \n", ct.Second())
		return
	}

	fmt.Println("Start():", "3, Starting...")

	// 激活冷却期
	cs.isCooldownActive = true
	cs.timeout = time.After(cs.cooldownDuration)
	/*
		// After waits for the duration to elapse and then sends the current time
		// on the returned channel.
		// It is equivalent to NewTimer(d).C.
		// The underlying Timer is not recovered by the garbage collector
		// until the timer fires. If efficiency is a concern, use NewTimer
		// instead and call Timer.Stop if the timer is no longer needed.
	*/
	go func() {
		fmt.Printf("Start(): 4, Started, will wait for %d nanosecond...\n", cs.cooldownDuration)
		<-cs.timeout
		fmt.Printf("Start(): 5, Stopping, has waited for %d nanosecond...\n", cs.cooldownDuration)
		fmt.Printf("Start(): 6, current time: %d \n", ct.UnixMilli())
		cs.isCooldownActive = false
		close(cs.cooldownCh) // 通知冷却期结束
	}()
}

func (cs *CooldownService) WaitForCooldown2() {
	fmt.Println("WaitForCooldown():", "0, ------------------------")
	if !cs.isCooldownActive {
		fmt.Println("WaitForCooldown():", "1, Not in cooldown, no need to wait.")
		return
	}

	fmt.Println("WaitForCooldown():", "2, Waiting for cooldown to end...")
	a := <-cs.cooldownCh // 阻塞直到冷却期结束
	fmt.Println(a)
	fmt.Println("WaitForCooldown():", "3, Cooldown ended, you can start again.")
}

func GoroutineChannelDemo9_2_cool_down_demo_1_start() {
	cooldownService := NewCooldownService(10 * time.Second)

	// 尝试启动，但不在冷却期，所以直接启动
	cooldownService.Start()
	// 假设我们在这里做了一些事情，然后想要再次启动，但此时处于冷却期
	fmt.Println("cool_down_demo_1_start():", "1, do a lot of things here!!")
	time.Sleep(2 * time.Second)

	cooldownService.Start()
	fmt.Println("cool_down_demo_1_start():", "2, do a lot of things here!!")
	time.Sleep(2 * time.Second)

	cooldownService.Start()
	fmt.Println("cool_down_demo_1_start():", "3, do a lot of things here!!")
	time.Sleep(2 * time.Second)

	// 冷却期结束后，我们可以再次尝试启动
	cooldownService.Start()
	time.Sleep(4 * time.Second)
	cooldownService.Start()

}

func GoroutineChannelDemo9_2_cool_down_demo_1_wait() {
	cooldownService := NewCooldownService(10 * time.Second)

	// 尝试启动，但不在冷却期，所以直接启动
	cooldownService.Start()

	// 假设我们在这里做了一些事情，然后想要再次启动，但此时处于冷却期
	fmt.Println("cool_down_demo_1_wait():", "1, do a lot of things here!!")
	time.Sleep(2 * time.Second)

	cooldownService.Start()
	fmt.Println("cool_down_demo_1_wait():", "2, do a lot of things here!!")
	time.Sleep(2 * time.Second)

	cooldownService.Start()
	fmt.Println("cool_down_demo_1_wait():", "3, do a lot of things here!!")
	time.Sleep(2 * time.Second)

	// 调用 WaitForCooldown 来等待冷却期结束
	fmt.Println("\ncool_down_demo_1_wait():", "4, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()

	fmt.Println("\ncool_down_demo_1_wait():", "5, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()

	// 冷却期结束后，我们可以再次尝试启动
	cooldownService.Start()
	time.Sleep(2 * time.Second)

	// 假设程序继续运行，并在某个时候再次需要等待冷却期结束
	fmt.Println("\ncool_down_demo_1_wait():", "6, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()

	fmt.Println("\ncool_down_demo_1_wait():", "7, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()

	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	fmt.Println("\ncool_down_demo_1_wait():", "8, cooldownService.WaitForCooldown()")
	cooldownService.WaitForCooldown()
	// 注意：在实际应用中，你可能需要处理更多的边界情况和错误
}

func GoroutineChannelDemo9_2_cool_down_demo_2_stop(cooldownService *CooldownService) {

	// 尝试启动，但不在冷却期，所以直接启动
	fmt.Println("在Channel_cool_down_demo_2_stop中，尝试启动，不在冷却期，所以直接启动")
	cooldownService.Start()
	fmt.Println("在Channel_cool_down_demo_2_stop中，尝试启动，不在冷却期，所以直接启动")

	//// 假设我们在这里做了一些事情，然后想要再次启动，但此时处于冷却期
	//fmt.Println("do a lot of things here!!")
	//time.Sleep(2 * time.Second)
	////fmt.Println("do a lot of things here!!")
	////time.Sleep(2 * time.Second)
	////fmt.Println("do a lot of things here!!")
	////time.Sleep(2 * time.Second)
	//
	//// 调用 WaitForCooldown 来等待冷却期结束
	//cooldownService.WaitForCooldown()
	//
	//for j := 0; j < 100; j++ {
	//	// 冷却期结束后，我们可以再次尝试启动
	//	cooldownService.Start()
	//}
	//
	//// 假设程序继续运行，并在某个时候再次需要等待冷却期结束
	//cooldownService.WaitForCooldown()
	//
	//// 注意：在实际应用中，你可能需要处理更多的边界情况和错误
}

func GoroutineChannelDemo9_2_cool_down_demo_2_start(cooldownService *CooldownService) {

	fmt.Println("在新的方法中尝试启动，但此时处于冷却期，所以不能启动")
	cooldownService.Start()
	fmt.Println("在新的方法中尝试启动，但此时处于冷却期，所以不能启动")

	// 假设我们在这里做了一些事情，然后想要再次启动，但此时处于冷却期
	fmt.Println("do a lot of things here!!")
	time.Sleep(2 * time.Second)
	//fmt.Println("do a lot of things here!!")
	//time.Sleep(2 * time.Second)
	//fmt.Println("do a lot of things here!!")
	//time.Sleep(2 * time.Second)

	//// 调用 WaitForCooldown 来等待冷却期结束
	//cooldownService.WaitForCooldown()

	for j := 0; j < 100; j++ {
		// 冷却期结束后，我们可以再次尝试启动
		//cooldownService.Start()
		if cooldownService.isCooldownActive {
			fmt.Println("Still within cool down period... ", j)
		} else {
			fmt.Println("Cool down period finished... ", j)
			break
		}

		time.Sleep(1 * time.Second)
	}

	fmt.Println("You can do other things now... ")

	//// 假设程序继续运行，并在某个时候再次需要等待冷却期结束
	//cooldownService.WaitForCooldown()

	// 注意：在实际应用中，你可能需要处理更多的边界情况和错误
}
