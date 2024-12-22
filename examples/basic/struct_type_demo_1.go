package basic

import "fmt"

// 定义一个结构体
type T struct {
	name string
}

// 一个方法的接收者是值类型
func (t T) method1() {
	fmt.Println("接收者是值类型")
	t.name = "new name1"
}

// 一个方法的接收者是指针类型
func (t *T) method2() {
	fmt.Println("接收者是指针类型")
	t.name = "new name2"
}

func TypeDemo1_value_receiver_vs_pointer_receiver() {

	t := T{"old name"}

	fmt.Println("method1 调用前 ", t.name)
	t.method1()
	fmt.Println("method1 调用后 ", t.name)

	fmt.Println("method2 调用前 ", t.name)
	t.method2()
	fmt.Println("method2 调用后 ", t.name)
}
