package xor_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/mrumyantsev/xor"
)

const aplhaLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func TestEncrypt(t *testing.T) {
	dataDefault := []byte(aplhaLetters)

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

	dataCopy := make([]byte, len(dataDefault))

	for _, keyEntry := range keysTestTable {
		for _, wnEntry := range workersTestTable {
			copy(dataCopy, dataDefault)

			// encryption + decryption should bring default data
			xor.EncryptData(dataCopy, keyEntry.key, wnEntry.encWNum)
			xor.EncryptData(dataCopy, keyEntry.key, wnEntry.decWNum)

			if !reflect.DeepEqual(dataCopy, dataDefault) {
				fmt.Println("decryption fail: default data does not match data copy")
				fmt.Println("def:", dataDefault)
				fmt.Println("cpy:", dataCopy)
				t.FailNow()
			}
		}
	}
}
