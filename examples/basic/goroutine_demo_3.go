package basic

import (
	"fmt"
	"time"
)

type fetchResult struct {
	Message string
	Error   error
}

func fetchA() fetchResult {
	time.Sleep(time.Second * 4)
	return fetchResult{"A data", nil}
}

func fetchB() fetchResult {
	time.Sleep(time.Second * 4)
	return fetchResult{"B data", nil}
}

func fetchC() fetchResult {
	time.Sleep(time.Second * 4)
	return fetchResult{"C data", nil}
}

func fetchConcurrent() {
	aChan := make(chan fetchResult, 0)
	bChan := make(chan fetchResult, 0)
	cChan := make(chan fetchResult, 0)

	go func(c chan fetchResult) {
		c <- fetchA()
	}(aChan)
	go func(c chan fetchResult) {
		c <- fetchB()
	}(bChan)
	go func(c chan fetchResult) {
		c <- fetchC()
	}(cChan)

	for i := 0; i < 3; i++ {
		select {
		case a := <-aChan:
			fmt.Println(a)
		case b := <-bChan:
			fmt.Println(b)
		case c := <-cChan:
			fmt.Println(c)

		}
	}
}
