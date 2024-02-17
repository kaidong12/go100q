package basic

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func Struct_demo() {
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book1 address:\n %p\n", &Book1)
	printBook1(Book1)

	/* 打印 Book2 信息 */
	fmt.Printf("Book2 address:\n %p\n", &Book2)
	printBook(&Book2)

	deferPanic()

}

func printBook1(book Books) {
	fmt.Printf("Book address:\n %p\n", &book)

	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
	fmt.Printf("All fields:\n %v\n", book)

}

func printBook(book *Books) {
	fmt.Printf("Book address:\n %p\n", book)

	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
	fmt.Printf("All fields:\n %v\n", book)

}

func deferPanic() {
	fmt.Println("deferPanic 1")
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	fmt.Println("deferPanic 2")
	panic("出错啦")
	fmt.Println("deferPanic 3")
}
