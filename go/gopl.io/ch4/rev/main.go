// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	type Protocol struct {
		Version     uint8
		BodyLen     uint16
		Reserved    [2]byte
		Unit        uint8
		Value       uint32
	}

	var p Protocol
	p.Version = 1
	p.BodyLen = 20
	p.Reserved = [2]byte{'a', 'b'}
	p.Unit = 3
	p.Value = 4

	buffer := new(bytes.Buffer)
	_ = binary.Write(buffer, binary.LittleEndian, p)

	bin := buffer.Bytes()
	fmt.Println("bin is", bin)

	var q Protocol
	_ = binary.Read(bytes.NewReader(bin), binary.LittleEndian, &q)
	fmt.Println("q is", q)

	fmt.Println(111111111111)

	a := "Hello, playground"
	fmt.Println([]byte(a))

	b := "Hello, playground"
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(strings.NewReader(b))
	fmt.Println(buf.Bytes())

	fmt.Println(222222222)

	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	line, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s", line)
	// 这里可以换上任意的 bufio 的 Read/Write 操作
	n, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line)
	fmt.Println(string(n))

	fmt.Println(333333333333)

	bs.Append([]byte{1, 2})
	fmt.Println("bs is", bs)

	fmt.Println(4444444444444)

	file, _ := os.Open("./test.txt")
	buf2 := bytes.NewBufferString("Hello world ")
	buf2.ReadFrom(file)              //将text.txt内容追加到缓冲器的尾部
	fmt.Println(buf2.String())


	var c bytes.Buffer       //直接定义一个Buffer变量，不用初始化，可以直接使用
	c1 := new(bytes.Buffer)   //使用New返回Buffer变量
	c2 := bytes.NewBuffer([]byte{1, 2})   //从一个[]byte切片，构造一个Buffer
	c3 := bytes.NewBufferString("asdf")   //从一个string变量，构造一个Buffer

	c.ReadByte()

	c.Write(d []byte)        //将切片d写入Buffer尾部
	c1.WriteString(s string)  //将字符串s写入Buffer尾部
	c2.WriteByte(c byte)     //将字符c写入Buffer尾部
	c.WriteRune(r rune)     //将一个rune类型的数据放到缓冲器的尾部
	c1.WriteTo(w io.Writer)  //将Buffer中的内容输出到实现了io.Writer接口的可写入对象中




}

type ByteSlice []byte
var bs ByteSlice

func (p *ByteSlice) Append(data []byte) {
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
	slice := *p
	// Body as above, without the return.
	slice = append(slice, data...)
	*p = slice
}

