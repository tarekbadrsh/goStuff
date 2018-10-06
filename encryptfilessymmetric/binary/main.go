package main

import (
	"flag"

	"github.com/tarekbadrshalaan/goStuff/encryptfilessymmetric"
)

// to encrypt
// go run main.go -f "filepath" -k yourkey
// to decrypt
// go run main.go -p d -f "filepath" -k yourkey

var f = flag.String("f", "file.txt", "your file path, defult 'file.txt'")
var k = flag.String("k", "foo", "your encryption key, defult 'foo'")
var p = flag.String("p", "e", "your operation {e:encrypt,d:decrypt}")

func main() {
	flag.Parse()
	if *p == "e" {
		encryptfilessymmetric.Encrypt(*f, *k)
	} else if *p == "d" {
		encryptfilessymmetric.Decrypt(*f, *k)
	}
}
