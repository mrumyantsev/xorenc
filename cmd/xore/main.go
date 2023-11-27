package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mrumyantsev/xore"
)

func main() {
	argsLength := len(os.Args)

	if argsLength < 3 {
		printUsageHelp()
		os.Exit(1)
	}

	filePath := os.Args[1]

	encryptKeyWords := os.Args[2:]
	encryptKeyChars := strings.Join(encryptKeyWords, " ")
	encryptKeyBytes := []byte(encryptKeyChars)

	err := xore.EcryptFile(filePath, encryptKeyBytes)
	if err != nil {
		fmt.Println("could not encrypt file. error: " + err.Error())
		os.Exit(2)
	}
}

func printUsageHelp() {
	fmt.Println("usage: xore <path/to/file> <any amount of words as encrypt key>")
}
