package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()
		_, err := fmt.Fprint(pw, "Hello")
		<-time.After(2 * time.Second)
		_, err = fmt.Fprintln(pw, "-World")
		if err != nil {
			panic(err)
		}
	}()

	_, err := io.Copy(os.Stdout, pr)

	if err != nil {
		panic(err)
	}
}
