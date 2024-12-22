package pointer

import "fmt"

// func swap(a int, b int) {
// 	fmt.Println(&a)
// 	fmt.Println(&b)

// 	var temp int
// 	temp = a
// 	a = b
// 	b = temp
// }

func swap(pa *int, pb *int) {
	fmt.Println(pa)
	fmt.Println(pb)
	fmt.Println(*pa)
	fmt.Println(*pb)

	var temp int
	temp = *pa //temp = main::a
	*pa = *pb  // main::a = main::b
	*pb = temp // main::b = temp

	fmt.Println(*pa)
	fmt.Println(*pb)
}

func Pointer_demo_1() {
	var a int = 10
	var b int = 20
	fmt.Println(&a)
	fmt.Println(&b)

	swap(&a, &b)
	// swap(a, b)

	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println("a = ", a, " b = ", b)

	var p *int

	p = &a

	fmt.Println(&a)
	fmt.Println(p)

	var pp **int //二级指针

	pp = &p

	fmt.Println(&p)
	fmt.Println(pp)
}
