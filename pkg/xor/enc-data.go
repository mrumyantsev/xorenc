package xor

import (
	"runtime"
	"sync"
)

const (
	// Default number of workers that used by data encrypter.
	NWorkersDefault = 0
)

var (
	nWorkers = 0
)

// Encrypt performs per-bit XOR encryption of data by the symbolic key.
//
// The decryption is provided by repeating of encryption with the same
// key, that was used for encryption.
//
// The data processing operation in this function is processed with the
// workers, which divides the data to the equal pieces and encrypts
// each. The number of workers is capped to the length of the data.
//
// If the number of workers is set to 0, then the number of workers
// will be equal to the number of physical processor cores.
func Encrypt(data []byte, key []byte) {
	var (
		dataLen = len(data)
		keyLen  = len(key)
	)

	if (dataLen == 0) || (keyLen == 0) {
		return
	}

	if nWorkers <= 0 {
		nWorkers = runtime.NumCPU()
	}

	if nWorkers > dataLen {
		nWorkers = dataLen
	}

	var (
		dataPart  = dataLen / nWorkers
		dataStart = 0
		dataEnd   = dataPart
		keyStart  = 0
		i         = 1 // loop counter
		wg        = &sync.WaitGroup{}
	)

	wg.Add(nWorkers)

	for {
		// Starting current worker.

		w := &worker{
			data:      data,
			key:       key,
			keyLen:    keyLen,
			keyStart:  keyStart,
			dataStart: dataStart,
			dataEnd:   dataEnd,
			wg:        wg,
		}

		go w.do()

		if i == nWorkers {
			break
		}

		i++

		// Calculating next worker data.

		dataStart += dataPart
		keyStart += dataPart

		for keyStart >= keyLen {
			keyStart -= keyLen
		}

		if i == nWorkers {
			dataEnd = dataLen
		} else {
			dataEnd += dataPart
		}
	}

	wg.Wait()
}

// NWorkers gets current number of workers that used to encrypt data
// inside Encrypt function.
func NWorkers() int {
	return nWorkers
}

// SetNWorkers sets how many workers will encrypt data inside Encrypt
// function.
func SetNWorkers(num int) {
	nWorkers = num
}

type worker struct {
	data      []byte
	key       []byte
	keyLen    int
	keyStart  int
	dataStart int
	dataEnd   int
	wg        *sync.WaitGroup
}

func (w *worker) do() {
	defer w.wg.Done()

	var (
		i = w.dataStart // data index
		j = w.keyStart  // key index
	)

	for i < w.dataEnd {
		if j >= w.keyLen {
			j = 0
		}

		w.data[i] = w.data[i] ^ w.key[j]

		i++
		j++
	}
}
