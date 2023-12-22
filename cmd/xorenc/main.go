package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mrumyantsev/xorenc"
)

const (
	_ERROR_EXIT_CODE int = 1
)

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		fmt.Println("usage: xorenc <path/to/file> <any number of words as an encryption key>")
		errorExit()
	}

	filePath := args[0]
	encryptKey := []byte(strings.Join(args[1:], " "))

	nbytes, err := xorenc.EncryptFile(filePath, encryptKey)
	if err != nil {
		fmt.Println(err.Error())
		errorExit()
	}

	fmt.Println("encrypted", nbytes, "bytes")
}

func errorExit() {
	os.Exit(_ERROR_EXIT_CODE)
}
