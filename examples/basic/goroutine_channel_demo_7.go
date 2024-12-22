package basic

import (
	"context"
	"fmt"
	"time"
)

func GoroutineChannelDemo7_WithTimeout_ctx_done() {
	readCtx, readCancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer func() {
		if readCancel != nil {
			readCancel()
		}
	}()

	times := 1
	for {
		select {
		case <-readCtx.Done():
			fmt.Println("Read operation completed or timed out")
			return
		default:
			fmt.Println("Reading data... ", times)
			times += 1
			time.Sleep(2 * time.Second)
		}
	}
	//
	//num := <-c //从c中接收数据，并赋值给num
	//fmt.Println("In main goroutine, 接收channel->c的数据，并赋值给num")
	//
	//fmt.Println("===========num = ", num)
	//fmt.Println("===========main go程结束")
}

func GoroutineChannelDemo7_WithTimeout_call_cancel() {
	readCtx, readCancel := context.WithTimeout(context.Background(), time.Duration(100)*time.Second)
	defer func() {
		if readCancel != nil {
			readCancel()
		}
	}()

	times := 1
	for {
		select {
		case <-readCtx.Done():
			fmt.Println("Read operation completed or timed out")
			return
		default:
			fmt.Println("Reading data... ", times)
			if times >= 5 {
				fmt.Println("Call readCancel() at times: ", times)
				readCancel()
				fmt.Println("readCancel() called!!!")
			}
			times += 1
			time.Sleep(1 * time.Second)
		}
	}
	//
	//num := <-c //从c中接收数据，并赋值给num
	//fmt.Println("In main goroutine, 接收channel->c的数据，并赋值给num")
	//
	//fmt.Println("===========num = ", num)
	//fmt.Println("===========main go程结束")
}
