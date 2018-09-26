package hashzeros

import (
	"encoding/hex"
	"strconv"
	"sync"
)

//GetBestZerosGoroutine :
func GetBestZerosGoroutine(text string, nonceLimit int, oprationLimit int) BestZero {
	var wg sync.WaitGroup
	data := []byte(text)

	finalBestZero := BestZero{}

	//!+handleHashs
	for i := 0; i < nonceLimit; i += oprationLimit {
		wg.Add(1)
		go func(text string, data []byte, start int, end int, wg *sync.WaitGroup) {
			defer wg.Done()
			currantBestZero := BestZero{}
			for x := start; x < end; x++ {
				c1 := HashTextNonce(data, x)
				if hasZeros, zerosResult := ChackHexadecimalZeros(c1[:], currantBestZero.Zeros); hasZeros {
					if len(currantBestZero.Zeros) < len(zerosResult) {
						currantBestZero = BestZero{text, x, text + strconv.Itoa(x), zerosResult, hex.EncodeToString(c1[:])}
					}
				}
			}

			if len(finalBestZero.Zeros) < len(currantBestZero.Zeros) {
				finalBestZero = currantBestZero
			}

		}(text, data, i, i+oprationLimit, &wg)
	}
	//!-handleHashs

	wg.Wait()
	return finalBestZero
}
