package xor

import (
	"sync"
)

// Encrypt performs per-bit XOR encryption of data by the key.
//
// The decryption is provided by repeating of encryption with the same
// key, that was used for encryption.
//
// The data processing operation in this function is processed with the
// workers, which divides the data to the equal pieces and encrypts
// each. The number of workers is capped to the length of the data, and
// it makes no processing, if given 0 or less.
func Encrypt(data []byte, key []byte, nWorkers int) {
	if (data == nil) || (key == nil) || (nWorkers <= 0) {
		return
	}

	var (
		dataLen int = len(data)
		keyLen  int = len(key)
	)

	if (dataLen == 0) || (keyLen == 0) {
		return
	}

	if nWorkers > dataLen {
		nWorkers = dataLen
	}

	var (
		dataPart  int = dataLen / nWorkers
		dataStart int = 0
		dataEnd   int = dataPart
		keyStart  int = 0
		i         int = 1 // loop counter

		wg sync.WaitGroup = sync.WaitGroup{}
	)

	wg.Add(nWorkers)

	for {
		// starting current worker

		go worker(workerData{
			data:      data,
			key:       key,
			keyStart:  keyStart,
			dataStart: dataStart,
			dataEnd:   dataEnd,
			wg:        &wg,
		})

		if i == nWorkers {
			break
		}

		i++

		// calculating next worker data

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

func worker(wd workerData) {
	var (
		keyLen  int = len(wd.key)
		dataLen int = wd.dataEnd
		i       int = wd.dataStart // data index
		j       int = wd.keyStart  // key index
	)

	for i < dataLen {
		if j >= keyLen {
			j = 0
		}

		wd.data[i] = wd.data[i] ^ wd.key[j]

		i++
		j++
	}

	wd.wg.Done()
}

type workerData struct {
	data      []byte
	key       []byte
	keyStart  int
	dataStart int
	dataEnd   int
	wg        *sync.WaitGroup
}
