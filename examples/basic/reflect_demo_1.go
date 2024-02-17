package basic

import (
	"fmt"
	"reflect"
)

func ReflectDemo_1() {
	var num float64 = 1.2345

	fmt.Println("type: ", reflect.TypeOf(num))
	fmt.Println("value: ", reflect.ValueOf(num))

}

//运行结果:
//type:  float64
//value:  1.2345

/*
	说明

	1. reflect.TypeOf： 直接给到了我们想要的type类型，如float64、int、各种pointer、struct 等等真实的类型
	2. reflect.ValueOf：直接给到了我们想要的具体的值，如1.2345这个具体数值，或者类似&{1 "Allen.Wu" 25} 这样的结构体struct的值
	3. 也就是说明反射可以将“接口类型变量”转换为“反射类型对象”，反射类型指的是reflect.Type和reflect.Value这两种
*/
