package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/mrumyantsev/xor"
)

const errorExitCode int = 1

func main() {
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

	nBytes, err = xor.EncryptFile(filePath, encryptKey, nCores)
	if err != nil {
		fmt.Println(err.Error())
		errorExit()
	}

	fmt.Println("encrypted", nBytes, "bytes")
}

func errorExit() {
	os.Exit(errorExitCode)
}
