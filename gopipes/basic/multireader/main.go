package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	header := strings.NewReader("<msg>")
	body := strings.NewReader("hello")
	footer := strings.NewReader("</msg>\n")

	readers := io.MultiReader(header, body, footer)

	_, err := io.Copy(os.Stdout, readers)
	if err != nil {
		panic(err)
	}
}
