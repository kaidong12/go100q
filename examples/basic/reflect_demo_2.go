package basic

import (
	"fmt"
	"reflect"
)

func ReflectDemo_with_convert() {
	var num float64 = 1.2345

	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	// 可以理解为“强制转换”，但是需要注意的时候，转换的时候，如果转换的类型不完全符合，则直接panic
	// Golang 对类型要求非常严格，类型一定要完全符合
	// 如下两个，一个是*float64，一个是float64，如果弄混，则会panic
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)

	fmt.Println(convertPointer)
	fmt.Println(convertValue)
}

//运行结果：
//0xc42000e238
//1.2345

/*
	说明

	1.  转换的时候，如果转换的类型不完全符合，则直接panic，类型要求非常严格！
	2.  转换的时候，要区分是指针还是指
	3.  也就是说反射可以将“反射类型对象”再重新转换为“接口类型变量”
*/
