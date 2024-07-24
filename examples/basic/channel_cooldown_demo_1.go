package basic

import (
	"fmt"
	"time"
)

type CooldownService struct {
	cooldownDuration time.Duration
	ticker           *time.Ticker
	//cooldownCh       chan struct{}
	isCooldownActive bool
}

func NewCooldownService(cooldownDuration time.Duration) *CooldownService {
	return &CooldownService{
		cooldownDuration: cooldownDuration,
		//cooldownCh:       make(chan struct{}),
	}
}

func (cs *CooldownService) Start() {
	var ct time.Time
	if cs.isCooldownActive {
		fmt.Println("Already in cooldown, please wait.")
		fmt.Println("current time: ", ct)
		return
	}

	fmt.Println("Starting...")
	// 模拟你的业务逻辑
	fmt.Println("do 你的业务逻辑 things here!!")

	// 激活冷却期
	cs.isCooldownActive = true
	cs.ticker = time.NewTicker(cs.cooldownDuration)
	go func() {
		ct = <-cs.ticker.C
		cs.ticker.Stop()
		cs.isCooldownActive = false
		//close(cs.cooldownCh) // 通知冷却期结束
	}()
}

//func (cs *CooldownService) WaitForCooldown() {
//	if !cs.isCooldownActive {
//		fmt.Println("Not in cooldown, no need to wait.")
//		return
//	}
//
//	fmt.Println("Waiting for cooldown to end...")
//	<-cs.cooldownCh // 阻塞直到冷却期结束
//	fmt.Println("Cooldown ended, you can start again.")
//}

func Channel_cool_down_demo_1() {
	cooldownService := NewCooldownService(10 * time.Second)

	// 尝试启动，但不在冷却期，所以直接启动
	cooldownService.Start()

	// 假设我们在这里做了一些事情，然后想要再次启动，但此时处于冷却期
	fmt.Println("do a lot of things here!!")
	time.Sleep(2 * time.Second)
	//fmt.Println("do a lot of things here!!")
	//time.Sleep(2 * time.Second)
	//fmt.Println("do a lot of things here!!")
	//time.Sleep(2 * time.Second)

	//// 调用 WaitForCooldown 来等待冷却期结束
	//cooldownService.WaitForCooldown()

	// 冷却期结束后，我们可以再次尝试启动
	cooldownService.Start()
	//
	//// 假设程序继续运行，并在某个时候再次需要等待冷却期结束
	//cooldownService.WaitForCooldown()

	// 注意：在实际应用中，你可能需要处理更多的边界情况和错误
}

func Channel_cool_down_demo_2_stop(cooldownService *CooldownService) {

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

func Channel_cool_down_demo_2_start(cooldownService *CooldownService) {

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
