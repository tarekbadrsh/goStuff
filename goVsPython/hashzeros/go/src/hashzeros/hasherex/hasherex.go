//time ./hasherex -t foo -c 1000000 > hasherex.trace
//go tool trace hasherex.trace

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"hashzeros"
	"os"
	"runtime/trace"
	"strconv"
	"strings"
	"sync"
)

//HashText :
func HashText(bytetext []byte) [32]byte {
	return sha256.Sum256(bytetext)
}

//GetBestZeros :
func GetBestZeros(text string, nonceLimit int, oprationLimit int) hashzeros.BestZero {
	data := []byte(text)
	zeros := "0"
	zerosCount := 1
	bestZero := hashzeros.BestZero{}

	wg := sync.WaitGroup{}
	for i := 0; i < nonceLimit; i += oprationLimit {
		wg.Add(1)
		go func(start int, end int) {
			defer wg.Done()
			for x := start; x < end; x++ {
				c1 := hashzeros.HashTextNonce(data, x)
				//c1 := HashTextNonce(data, x)
				if hasZeros, zerosinString := hashzeros.ChackHexadecimalZeros(c1[:], zeros); hasZeros {
					if zerosCount < len(zerosinString) {
						zerosCount = len(zerosinString)
						zeros = strings.Repeat("0", zerosCount)
						bestZero.Text = text + strconv.Itoa(x)
						bestZero.Nonce = x
						bestZero.Checksum = hex.EncodeToString(c1[:])
					}
				}
			}
		}(i, i+oprationLimit)
	}
	wg.Wait()

	return bestZero
}

var tex = flag.String("t", " ", "text need to hash")
var coun = flag.Int("c", 1000000, "count time hash")

func main() {
	trace.Start(os.Stdout)
	defer trace.Stop()

	flag.Parse()
	text := *tex
	count := *coun

	// data := []byte(text)
	// c1 := HashText(data)
	// res := hex.EncodeToString(c1[:])
	// fmt.Println(res)

	// pkgResult := hashzeros.GetBestZeros(text, count, 1000)
	// fmt.Println(pkgResult)

	_ = hashzeros.GetBestZeros(text, count, 1000)
}
