package main

import (
	_ "bytes"
	"io"
	"os"
	_ "time"
)

type Reader interface {
	Reads(p []byte) (n int, err error)
}

var a rune
var w io.Writer = os.Stdout // OK: *os.File 实现了Write方法
//var w io.Writer = new(bytes.Buffer) // OK: *bytes.Buffer 实现了Write方法
//var w io.Writer = time.Second       // 编译错误: 缺少Write方法, 未实现io.Writer接口

var rwc io.ReadWriteCloser = os.Stdout // OK: *os.File 实现了 Read, Write, Close方法
//var rwc io.ReadWriteCloser = new(bytes.Buffer) // 编译错误: *bytes.Buffer 缺少Close方法

//w = rwc  // OK: io.ReadWriteCloser实现了Write方法
//rwc = w  // 编译错误: io.Writer未实现Close方法
