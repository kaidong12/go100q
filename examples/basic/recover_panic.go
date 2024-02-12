package basic

import "fmt"

func divide(a, b int) {
	if b == 0 {
		panic("division by zero")
	}
	fmt.Println(a / b)
}

func recover_demo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	divide(10, 5)
	divide(10, 0)
	divide(10, 2)
	fmt.Println("End of program")
}
