package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var buffPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("Allocating new bytes.Buffer")
		return new(bytes.Buffer)
	},
}

func log1(w io.Writer, msg string) {
	b := buffPool.Get().(*bytes.Buffer)

	b.Reset()
	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(msg)
	b.WriteString("\n")
	_, err := w.Write(b.Bytes())
	if err != nil {
		return
	}

	buffPool.Put(b)
}

func main() {
	log1(os.Stdout, "first message")
	log1(os.Stdout, "second message")
}
