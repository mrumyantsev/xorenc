package main

import (
	"fmt"
	"log"

	"github.com/mrumyantsev/xore"
)

const (
	_FILE_TO_ENCRYPT = "Daisy.txt"
)

var (
	encryptKey = []byte("SuNnY DaY")
)

func main() {
	err := xore.EcryptFile(_FILE_TO_ENCRYPT, encryptKey)
	if err != nil {
		log.Fatal("could not encrypt file. error:", err)
	}

	fmt.Println(_FILE_TO_ENCRYPT, "encrypted")
}
