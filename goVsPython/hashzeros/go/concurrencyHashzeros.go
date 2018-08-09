package hashzeros

import (
	"encoding/hex"
	"strconv"
	"sync"
)

//!+broadcaster

//Concurrencybroadcaster :
func Concurrencybroadcaster(handleHashsResult *chan BestZero, processResult *chan BestZero) {
	currantBestZero := BestZero{}
	for {
		select {
		case res := <-*handleHashsResult:
			if len(currantBestZero.Zeros) < len(res.Zeros) {
				currantBestZero = res
			}
		case <-*processResult:
			*processResult <- currantBestZero
		}
	}
}

//!-broadcaster

//!+handleHashs

//ConcurrencyhandleHashs :
func ConcurrencyhandleHashs(text string, data []byte, start int, end int, wg *sync.WaitGroup, handleHashsResult *chan BestZero) {
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
	*handleHashsResult <- currantBestZero
}

//!-handleHashs

//ConcurrencyGetBestZeros :
func ConcurrencyGetBestZeros(text string, nonceLimit int, oprationLimit int) BestZero {
	//channels
	var handleHashsResult = make(chan BestZero)
	var processResult = make(chan BestZero)

	var wg sync.WaitGroup
	data := []byte(text)

	//!+broadcaster
	go Concurrencybroadcaster(&handleHashsResult, &processResult)
	//!-broadcaster

	//!+handleHashs
	for i := 0; i < nonceLimit; i += oprationLimit {
		wg.Add(1)
		go ConcurrencyhandleHashs(text, data, i, i+oprationLimit, &wg, &handleHashsResult)
	}

	//!-handleHashs

	wg.Wait()
	processResult <- BestZero{}
	return <-processResult
}
