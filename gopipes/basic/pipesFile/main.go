package main

import (
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "missing paths of images to cat")
		os.Exit(2)
	}

	for _, path := range os.Args[1:] {
		if err := cat(path); err != nil {
			fmt.Fprintf(os.Stderr, "could not cat %s: %v\n", path, err)
		}
	}
}
func cat(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "could not open image")
	}
	defer file.Close()
	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()
		if _, err = io.Copy(pw, file); err != nil {
			pw.CloseWithError(errors.Wrap(err, "could not copy the image"))
			return
		}
	}()

	newImage, err := os.Create("newImage.png")
	if err != nil {
		return errors.Wrap(err, "could not create new image")
	}
	_, err = io.Copy(newImage, pr)
	if err != nil {
		return errors.Wrap(err, "could not copy to create new image")
	}
	return nil
}
