package basic

import (
	"fmt"
	"io"
	"os"
)

/*
在Golang的实现中，每个interface变量都有一个对应pair，pair中记录了实际变量的值和类型:
(value, type)

value是实际变量值，type是实际变量的类型。
一个interface{}类型的变量包含了2个指针，一个指针指向值的类型【对应concrete type】，另外一个指针指向实际的值【对应value】。

*/

func ReflectPair_pass_1() {

	// 创建类型为*os.File的变量，然后将其赋给一个接口变量r：
	tty, err := os.OpenFile("testdata/output.txt", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	// 接口变量r的pair中将记录如下信息：(tty, *os.File)，这个pair在接口变量的连续赋值过程中是不变的，将接口变量r赋给另一个接口变量w:
	var r io.Reader
	r = tty

	// 接口变量w的pair与r的pair相同，都是:(tty, *os.File)，即使w是空接口类型，pair也是不变的。
	var w io.Writer
	w = r.(io.Writer)

	_, err = w.Write([]byte("HELLO THIS IS A TEST!!!\n"))
	if err != nil {
		return
	}

}
