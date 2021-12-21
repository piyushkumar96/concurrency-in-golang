package main

import (
	"bytes"
	"io"
	"os"
	"time"
)

func log(w io.Writer, msg string) {
	var b bytes.Buffer

	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(msg)
	b.WriteString("\n")
	_, err := w.Write(b.Bytes())
	if err != nil {
		return
	}

}

func main() {
	log(os.Stdout, "first message")
	log(os.Stdout, "second message")
}
