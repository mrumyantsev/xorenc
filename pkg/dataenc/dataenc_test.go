package dataenc

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncrypt(t *testing.T) {
	defaultData := []byte("1n@kjnd lk#j34 h2@bg$7 6hb%3kjb&d8s d?8 a9s#dy#97")

	testTable := []struct {
		key []byte
	}{
		{key: []byte(" ")},
		{key: []byte("1111")},
		{key: []byte("v2@sV5n#^q4Z#i$g7x")},
	}

	dataEnc := New()
	dataCopy := make([]byte, len(defaultData))
	var isDataMatch bool

	for _, entry := range testTable {
		copy(dataCopy, defaultData)

		dataEnc.Encrypt(dataCopy, entry.key) // encryption
		dataEnc.Encrypt(dataCopy, entry.key) // decryption

		isDataMatch = reflect.DeepEqual(dataCopy, defaultData)

		if !isDataMatch {
			fmt.Println("decryption fail: default data does not match data copy")
			fmt.Println("def", defaultData)
			fmt.Println("cpy", dataCopy)
			t.FailNow()
		}
	}
}
