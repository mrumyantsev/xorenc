package xor

import (
	"flag"
	"fmt"
	"os"
	"strings"

	fileops "github.com/mrumyantsev/xor/internal/pkg/file-ops"
)

const (
	errorExitCode int = 1
)

type Xor struct {
}

func New() *Xor {
	return &Xor{}
}

func (x *Xor) Run() {
	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		fmt.Println("usage: xor <path/to/file> <any number of words as an encryption key>")
		errorExit()
	}

	filePath := args[0]
	encryptKey := []byte(strings.Join(args[1:], " "))

	nbytes, err := fileops.EncryptFile(filePath, encryptKey)
	if err != nil {
		fmt.Println(err.Error())
		errorExit()
	}

	fmt.Println("encrypted", nbytes, "bytes")
}

func errorExit() {
	os.Exit(errorExitCode)
}
