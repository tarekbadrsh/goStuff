package main

import (
	"flag"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/tarekbadrshalaan/goStuff/encryptfilessymmetric"
)

// to encrypt
// go run main.go -in in.txt -out decryptfile
// to decrypt
// go run main.go -in decryptfile -out rowfile

var in = flag.String("in", "in.txt", "your input file path, default 'in.txt'")
var out = flag.String("out", "out.txt", "your output file path, default 'out.txt'")

func main() {
	flag.Parse()
	operation := promptui.Select{
		Label: "your operation",
		Items: []string{"encrypt", "decrypt"},
	}
	_, result, err := operation.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	prompt := promptui.Prompt{
		Label: "your encryption password",
		Mask:  1, // allows hiding private information like password.
	}

	key, err := prompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	if result == "encrypt" {
		encryptfilessymmetric.EncryptFile(*in, *out, key)
	} else if result == "decrypt" {
		encryptfilessymmetric.DecryptFile(*in, *out, key)
	}
}
