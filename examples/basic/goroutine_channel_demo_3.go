package basic

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func GoroutineChannelDemo3_buffered_int() {
	c := make(chan int, 3) //带缓冲的通道

	//内置函数 len 返回未被读取的缓冲元素数量， cap 返回缓冲区大小
	fmt.Printf("len(c)=%d, cap(c)=%d\n", len(c), cap(c))

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 5; i++ {
			fmt.Printf("1, In sub goroutine, 发送--%d--到channel->c\n", i)
			c <- i
			fmt.Printf("4, 子go程正在运行[%d]: len(c)=%d, cap(c)=%d\n", i, len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second) //延时2s
	for i := 0; i < 5; i++ {
		fmt.Println("2, In main goroutine, 接收channel->c的数据，并赋值给num")
		num := <-c //从c中接收数据，并赋值给num
		fmt.Println("3, ===========num = ", num, "\n")
		time.Sleep(1 * time.Second)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("main进程结束")
}

func GoroutineChannelDemo3_slide_window1() {
	reader, err1 := os.Open("testdata/trace_ocu4.raw")
	if err1 != nil {
		fmt.Println("err1", err1)
		return
	}
	defer reader.Close()

	markerMap := make(map[byte]bool)

	var marker = []byte{0x80, 0x4d, 0x4c, 0x50} //{128, 77, 76, 80}
	//var marker = []byte{0x49, 0x6e, 0x42, 0x6f}

	for _, b := range marker {
		markerMap[b] = true
	}
	//fmt.Println(markerMap)

	window := make([]byte, len(marker))
	_, err2 := io.ReadFull(reader, window)
	if err2 != nil {
		fmt.Println("err2", err2)
		return
	}
	//fmt.Println(window)

	buf := make([]byte, 1)
	for {
		if !markerMap[window[3]] { // no match, move a window
			//fmt.Println(window)
			_, err3 := io.ReadFull(reader, window)
			if err3 != nil {
				fmt.Println("err3", err3)
				return
			}
		} else { // has match, move a byte
			_, err4 := reader.Read(buf)
			//fmt.Println(buf)
			if err4 != nil {
				fmt.Println("err4", err4)
				return
			}
			copy(window[:3], window[1:])
			window[3] = buf[0]
			//fmt.Println(window)
		}
		//fmt.Println(window)

		// found header start, return to parse header
		if bytes.Equal(window, marker) {
			//fmt.Println(window)
			//return
		}
	}
}

func GoroutineChannelDemo3_slide_window2() {
	var pattern = []byte{0x80, 0x4d, 0x4c, 0x50} //{128, 77, 76, 80}

	matches := []int{}
	window := make([]byte, 0)

	reader, err1 := os.Open("testdata/trace_ocu4.raw")
	if err1 != nil {
		fmt.Println("err1", err1)
		return
	}

	r := bufio.NewReader(reader)
	//buf := make([]byte, 1)
	for {
		// Read a line from the file
		line, _, err := r.ReadLine()
		//line, err := r.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading file:", err)
			}
			break
		}

		// Extract bytes from the line
		bytesRead := len(line)
		window = append(window, line...)

		// Slide the window and search for the pattern
		for i := 0; i+len(pattern) <= bytesRead; i++ {
			if bytes.Equal(window[i:i+len(pattern)], pattern) {
				matches = append(matches, i)
			}
		}

		// Remove bytes from the front of the window
		window = window[bytesRead:]
	}

	fmt.Println(matches)
}

func GoroutineChannelDemo3_slide_window3() {
	var pattern = []byte{0x80, 0x4d, 0x4c, 0x50} //{128, 77, 76, 80}

	reader, err := os.Open("testdata/trace_ocu4.raw")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	start, end := 0, 0
	buf := make([]byte, 1)

	for {
		_, err2 := io.ReadFull(reader, buf)
		if err2 != nil {
			fmt.Println("err3", err2)
			return
		}

		if buf[0] == pattern[0] {
			match := true
			for i := 1; i < len(pattern); i++ {
				_, err3 := io.ReadFull(reader, buf)
				if err3 != nil {
					fmt.Println("err3", err3)
					return
				}

				if buf[0] != pattern[i] {
					match = false
					break
				}
			}

			if match {
				//return
				//fmt.Println(pattern)
			}
		}

		end++
		if end-start >= len(pattern) {
			start++
		}
	}
}
