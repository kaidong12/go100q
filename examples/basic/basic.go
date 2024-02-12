package basic

import (
	"fmt"
	"math"
)

func Variable() {
	fmt.Printf("Hello, World!\n")
	b := 111
	var s string = fmt.Sprintf("%s%d%d", "sss", b, 222)
	fmt.Printf(s)
}

func Array_demo() {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
	fmt.Println()

	a1 := [5]int{1, 2}
	for j = 0; j < cap(a1); j++ {
		fmt.Printf("Element[%d] = %d\n", j, a1[j])
	}
	fmt.Println()

	a2 := [5]int{2: 1, 3: 2, 4: 3}
	for j = 0; j < cap(a2); j++ {
		fmt.Printf("Element[%d] = %d\n", j, a2[j])
	}
	fmt.Println()

	a3 := [...]int{1, 2, 3, 4, 5}
	for j = 0; j < cap(a2); j++ {
		fmt.Printf("Element[%d] = %d\n", j, a3[j])
	}
	fmt.Println()

}

func Slice_and_range_demo() {

	var nums = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range nums {
		fmt.Printf("%d ** %d = %d \n", 2, i+1, int(math.Pow(2, v)))
	}

	// s := make([]int,len,cap)
	numbers := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers), cap(numbers), numbers)

	numbers1 := nums[:8]
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers1), cap(numbers1), numbers1)

	numbers2 := nums[2:8]
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers2), cap(numbers2), numbers2)

	numbers3 := nums[2:]
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers3), cap(numbers3), numbers3)

}

func Map_demo() {

	// var siteMap map[string]string
	// siteMap = make(map[string]string)

	ages := map[string]int{
		"Joe":  30,
		"Mary": 25,
		"John": 35,
	}
	ages["John"] = 30
	ages["Mark"] = 28
	ages["Jane"] = 29

	fmt.Println()
	fmt.Println(ages)
	fmt.Println("John's age: ", ages["John"])

	if age, ok := ages["bob"]; !ok {
		fmt.Printf("bob is not exist %d\n", age)
	}

	age, ok := ages["bob"]
	fmt.Println(age)
	if !ok {
		fmt.Println("bob is not exist")
	}

	fmt.Println()
	siteMap := make(map[string]string)

	/* map 插入 key - value 对,各个国家对应的首都 */
	siteMap["Google"] = "谷歌"
	siteMap["Runoob"] = "菜鸟教程"
	siteMap["Baidu"] = "百度"
	siteMap["Wiki"] = "维基百科"
	siteMap["Facebook"] = "脸书"

	for site := range siteMap {
		fmt.Println(site, "站点是", siteMap[site])
	}

	/*查看元素在集合中是否存在 */
	name, ok := siteMap["Facebook"] /*如果确定是真实的,则存在,否则不存在 */
	fmt.Println(ok)
	fmt.Println(name)

	if ok {
		fmt.Println("Facebook 的站点是", name)
	} else {
		fmt.Println("Facebook 站点不存在")
	}

}

func Recursion_demo() {

	for i := 1; i <= 10; i++ {
		fmt.Println(fibonacci(i))
	}
}

func fibonacci(n int) int {

	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func Print_prime(n int) {

	var i, j int
	var flag bool
	for i = 2; i <= n; i++ {
		flag = true
		for j = 2; j < (i/2 + 1); j++ {

			if i%j == 0 {
				flag = false
				break
			}
		}

		if flag == true {
			fmt.Printf("%d \n", i)
		}
	}
}

func Pointer() {

	var a int = 20
	var a_pointer *int

	a_pointer = &a

	fmt.Printf("a的地址: %x\n", &a)
	fmt.Printf("a_pointer的值: %x\n", a_pointer)
	fmt.Printf("a_pointer指向的值: %d\n", *a_pointer)
}
