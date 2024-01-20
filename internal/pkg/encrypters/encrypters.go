package encrypters

import (
	"github.com/mrumyantsev/xor/pkg/dataenc"
	"github.com/mrumyantsev/xor/pkg/fileenc"
)

type DataEncrypter interface {
	Encrypt(data []byte, key []byte) (nBytes int)
}

type FileEncrypter interface {
	Encrypt(path string, key []byte) (nBytes int, err error)
}

type Encrypters struct {
	DataEncrypter
	FileEncrypter
}

func New(dataEnc *dataenc.DataEnc) *Encrypters {
	return &Encrypters{
		DataEncrypter: dataenc.New(),
		FileEncrypter: fileenc.New(dataEnc),
	}
}
