package basic

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	"log/slog"
	"net"
	"os"
	"strings"
)

/*

package io

// Package io provides basic interfaces to I/O primitives.
// Its primary job is to wrap existing implementations of such primitives,
// such as those in package os, into shared public interfaces that
// abstract the functionality, plus some other related primitives.
//
// Because these interfaces and primitives wrap lower-level operations with
// various implementations, unless otherwise informed clients should not
// assume they are safe for parallel execution.


type Reader interface {
	Read(p []byte) (n int, err error)
}

// Reader is the interface that wraps the basic Read method.
//
// Read reads up to len(p) bytes into p. It returns the number of bytes
// read (0 <= n <= len(p)) and any error encountered. Even if Read
// returns n < len(p), it may use all of p as scratch space during the call.
// If some data is available but not len(p) bytes, Read conventionally
// returns what is available instead of waiting for more.
//
// When Read encounters an error or end-of-file condition after
// successfully reading n > 0 bytes, it returns the number of
// bytes read. It may return the (non-nil) error from the same call
// or return the error (and n == 0) from a subsequent call.
// An instance of this general case is that a Reader returning
// a non-zero number of bytes at the end of the input stream may
// return either err == EOF or err == nil. The next Read should
// return 0, EOF.
//
// Callers should always process the n > 0 bytes returned before
// considering the error err. Doing so correctly handles I/O errors
// that happen after reading some bytes and also both of the
// allowed EOF behaviors.
//
// If len(p) == 0, Read should always return n == 0. It may return a
// non-nil error if some error condition is known, such as EOF.
//
// Implementations of Read are discouraged from returning a
// zero byte count with a nil error, except when len(p) == 0.
// Callers should treat a return of 0 and nil as indicating that
// nothing happened; in particular it does not indicate EOF.
//
// Implementations must not retain p.
*/

/*
package strings

A Reader implements the [io.Reader], [io.ReaderAt], [io.ByteReader], [io.ByteScanner],
[io.RuneReader], [io.RuneScanner], [io.Seeker], and [io.WriterTo] interfaces by reading
from a string.
The zero value for Reader operates like a Reader of an empty string.

	type Reader struct {
		s        string
		i        int64 // current reading index
		prevRune int   // index of previous rune; or < 0
	}
*/

func IO_reader_demo_1() {
	reader := strings.NewReader("NewReader returns a new [Reader] reading from s.\n It is similar to [bytes.NewBufferString] but more efficient and non-writable.")
	buffer := make([]byte, 8)
	for {
		n, err := reader.Read(buffer)
		fmt.Printf("n = %v, err = %v, b = %v \n", n, err, buffer)
		fmt.Printf("b[:n] = %q\n", buffer[:n])
		if err == io.EOF {
			break
		}
	}
}

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	if b >= 'A' && b <= 'M' || b >= 'a' && b <= 'm' {
		return b + 13
	}

	if b >= 'N' && b <= 'Z' || b >= 'n' && b <= 'z' {
		return b - 13
	}

	return b
}

func (r13r rot13Reader) Read(b []byte) (int, error) {
	n, err := r13r.r.Read(b)
	for i, v := range b {
		b[i] = rot13(v)
	}

	return n, err
}

func IO_reader_demo_2() {
	reader := strings.NewReader("NewReader returns a new [Reader] reading from s.\n It is similar to [bytes.NewBufferString] but more efficient and non-writable.")
	r := rot13Reader{reader}
	_, err := io.Copy(os.Stdout, &r)
	if err != nil {
		fmt.Println(err)
	}
}

func IO_reader_demo_3() {
	input := []byte("foo\x00bar")
	buffer := new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.StdEncoding, buffer)
	encoder.Write(input)
	fmt.Println(string(buffer.Bytes()))
}

type Protocol struct {
	Version  uint8
	BodyLen  uint16
	Reserved [2]byte
	Unit     uint8
	Value    uint32
}

func IO_reader_demo_4_read() {
	var p Protocol
	var bin []byte
	//...
	binary.Read(bytes.NewReader(bin), binary.LittleEndian, &p)
}

func IO_reader_demo_4_write() {
	var p Protocol
	buffer := new(bytes.Buffer)
	//...
	binary.Write(buffer, binary.LittleEndian, p)
	bin := buffer.Bytes()
	fmt.Printf("%q", bin)
}

/*
package bufio

// Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer
// object, creating another object (Reader or Writer) that also implements
// the interface but provides buffering and some help for textual I/O.

// Reader implements buffering for an io.Reader object.
type Reader struct {
	buf          []byte
	rd           io.Reader // reader provided by the client
	r, w         int       // buf read and write positions
	err          error
	lastByte     int // last byte read for UnreadByte; -1 means invalid
	lastRuneSize int // size of last rune read for UnreadRune; -1 means invalid
}


// Read reads data into p.
// It returns the number of bytes read into p.
// The bytes are taken from at most one Read on the underlying [Reader],
// hence n may be less than len(p).
// To read exactly len(p) bytes, use io.ReadFull(b, p).
// If the underlying [Reader] can return a non-zero count with io.EOF,
// then this Read method can do so as well; see the [io.Reader] docs.
func (b *Reader) Read(p []byte) (n int, err error) {
	n = len(p)
	if n == 0 {
		if b.Buffered() > 0 {
			return 0, nil
		}
		return 0, b.readErr()
	}
	if b.r == b.w {
		if b.err != nil {
			return 0, b.readErr()
		}
		if len(p) >= len(b.buf) {
			// Large read, empty buffer.
			// Read directly into p to avoid copy.
			n, b.err = b.rd.Read(p)
			if n < 0 {
				panic(errNegativeRead)
			}
			if n > 0 {
				b.lastByte = int(p[n-1])
				b.lastRuneSize = -1
			}
			return n, b.readErr()
		}
		// One read.
		// Do not use b.fill, which will loop.
		b.r = 0
		b.w = 0
		n, b.err = b.rd.Read(b.buf)
		if n < 0 {
			panic(errNegativeRead)
		}
		if n == 0 {
			return 0, b.readErr()
		}
		b.w += n
	}

	// copy as much as we can
	// Note: if the slice panics here, it is probably because
	// the underlying reader returned a bad count. See issue 49795.
	n = copy(p, b.buf[b.r:b.w])
	b.r += n
	b.lastByte = int(b.buf[b.r-1])
	b.lastRuneSize = -1
	return n, nil
}

*/

func IO_reader_demo_5(ctx *gin.Context) {
	var conn net.Conn
	upGrader := websocket.Upgrader{}
	wsconn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		slog.Error("can't upgrade req to ws.", "err_info", err.Error())
		return
	}
	//...
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err == nil {
			writeErr := wsconn.WriteJSON(line[:len(line)-1])
			if writeErr != nil {
				break
			}
			continue
		}
	}
}
