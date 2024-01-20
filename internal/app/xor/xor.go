package xor

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mrumyantsev/xor/internal/pkg/encrypters"
	"github.com/mrumyantsev/xor/pkg/dataenc"
)

const (
	errorExitCode int = 1
)

type Xor struct {
	encrypters *encrypters.Encrypters
}

func New() *Xor {
	dataEnc := dataenc.New()

	encrypters := encrypters.New(dataEnc)

	return &Xor{encrypters: encrypters}
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

	nbytes, err := x.encrypters.FileEncrypter.Encrypt(filePath, encryptKey)
	if err != nil {
		fmt.Println(err.Error())
		errorExit()
	}

	fmt.Println("encrypted", nbytes, "bytes")
}

func errorExit() {
	os.Exit(errorExitCode)
}
