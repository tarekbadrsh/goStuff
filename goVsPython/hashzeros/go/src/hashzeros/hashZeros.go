package hashzeros

import (
	"encoding/hex"
	"strconv"
	"sync"
)

//channels
var handleHashsResult = make(chan BestZero)
var endProcess = make(chan bool)
var processResult = make(chan BestZero)

//!+broadcaster
func broadcaster() {
	currantBestZero := BestZero{}
	for {
		select {
		case res := <-handleHashsResult:
			if len(currantBestZero.Zeros) < len(res.Zeros) {
				currantBestZero = res
			}
		case end := <-endProcess:
			if end {
				processResult <- currantBestZero
			}
		}
	}

}

//!-broadcaster

//!+handleHashs

func handleHashs(text string, data []byte, start int, end int, wg *sync.WaitGroup) {
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
	handleHashsResult <- currantBestZero
}

//!-handleHashs

//GetBestZeros :
func GetBestZeros(text string, nonceLimit int, oprationLimit int) BestZero {
	var wg sync.WaitGroup
	data := []byte(text)

	go broadcaster()
	for i := 0; i < nonceLimit; i += oprationLimit {
		wg.Add(1)
		go handleHashs(text, data, i, i+oprationLimit, &wg)
	}
	wg.Wait()
	endProcess <- true
	return <-processResult
}
