package xor_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/mrumyantsev/xor/pkg/xor"
)

const (
	aplhaLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func TestEncrypt(t *testing.T) {
	dataOrig := []byte(aplhaLetters)

	keysTestTable := []struct {
		key []byte
	}{
		{[]byte(" ")},
		{[]byte("abc")},
		{[]byte("1234")},
		{[]byte("v2@sV5n#^q4Z#i$g7x")},
		{[]byte(aplhaLetters)},
	}

	workersTestTable := []struct {
		encWNum int
		decWNum int
	}{
		{1, 2},
		{4, 5},
		{7, 4},
		{9, 5},
		{13, 12},
		{10_000, 100_000},
	}

	dataCopy := make([]byte, len(dataOrig))

	for _, keyEntry := range keysTestTable {
		for _, wnEntry := range workersTestTable {
			copy(dataCopy, dataOrig)

			fmt.Println("orig:", dataOrig)

			// Provide encryption and decryption that restores original
			// data.

			xor.SetNWorkers(wnEntry.encWNum)

			xor.Encrypt(dataCopy, keyEntry.key)

			fmt.Println("copy:", dataCopy)

			xor.SetNWorkers(wnEntry.decWNum)

			xor.Encrypt(dataCopy, keyEntry.key)

			if !reflect.DeepEqual(dataCopy, dataOrig) {
				fmt.Println("decryption fail: original data does not match data copy")
				fmt.Println("orig:", dataOrig)
				fmt.Println("copy:", dataCopy)

				t.FailNow()
			}
		}
	}
}
