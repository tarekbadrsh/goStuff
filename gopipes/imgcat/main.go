package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("missing path of image to cat")
	}

	for _, path := range os.Args[1:] {
		if err := cat(path); err != nil {
			fmt.Fprintf(os.Stderr, "could not cat %s: %v\n", path, err)
		}
	}
}

func cat(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "Could not open image")
	}
	defer f.Close()

	return nil
}
