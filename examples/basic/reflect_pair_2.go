package basic

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体类型
type Book struct {
	name string
}

func (this *Book) ReadBook() {
	fmt.Println("Read a book:", this.name)
}

func (this *Book) WriteBook() {
	fmt.Println("Write a book:", this.name)
}

func ReflectPair_pass_2() {
	b := &Book{"Golang study"}

	var r Reader
	r = b

	r.ReadBook()

	var w Writer
	w = r.(Writer)
	w.WriteBook()

}
