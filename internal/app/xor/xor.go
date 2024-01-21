package xor

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/mrumyantsev/xor/internal/pkg/encrypters"
	"github.com/mrumyantsev/xor/pkg/dataenc"
)

const errorExitCode int = 1

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

	var args []string = flag.Args()

	if len(args) < 2 {
		fmt.Println("usage: xor <path/to/file> <any number of words as an encryption key>")
		errorExit()
	}

	var (
		filePath   string = args[0]
		encryptKey []byte = []byte(strings.Join(args[1:], " "))
		nCores     int    = runtime.NumCPU()
		nBytes     int    = 0
		err        error  = nil
	)

	nBytes, err = x.encrypters.FileEncrypter.Encrypt(filePath, encryptKey, nCores)
	if err != nil {
		fmt.Println(err.Error())
		errorExit()
	}

	fmt.Println("encrypted", nBytes, "bytes")
}

func errorExit() {
	os.Exit(errorExitCode)
}
