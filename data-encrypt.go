package xor

import "sync"

// Encrypt performs per-bit XOR encryption of data by the key.
//
// The decryption is provided by repeating of encryption with the same
// key, that was in use in encryption.
//
// The data processing operation in this function is processed with the
// workers, which divides the data to the equal pieces and encrypts
// each. The number of workers is capped to the length of the data, and
// it makes no processing, if given 0 or less.
//
// Returns the number of the encrypted bytes.
func EncryptData(data []byte, key []byte, nWorkers int) (nBytes int) {
	if nWorkers <= 0 {
		return 0
	}

	var dataLen int = len(data)

	if nWorkers > dataLen {
		nWorkers = dataLen
	}

	var (
		keyLen    int = len(key)
		dataPart  int = dataLen / nWorkers
		dataStart int = 0
		dataEnd   int = dataPart
		keyStart  int = 0

		wg sync.WaitGroup = sync.WaitGroup{}
	)

	wg.Add(nWorkers)

	for i := int(0); i < nWorkers; i++ {
		if i == (nWorkers - 1) {
			dataEnd = dataLen
		}

		go worker(workerData{
			data:      data,
			key:       key,
			keyStart:  keyStart,
			dataStart: dataStart,
			dataEnd:   dataEnd,
			wg:        &wg,
		})

		dataStart += dataPart
		dataEnd += dataPart
		keyStart += dataPart

		for keyStart >= keyLen {
			keyStart -= keyLen
		}
	}

	wg.Wait()

	return dataLen
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
