package hashzeros

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

//HashTextNonce :
func HashTextNonce(bytetext []byte, nonce int) [32]byte {
	dataToHash := append(bytetext, []byte(strconv.Itoa(nonce))...)
	return sha256.Sum256(dataToHash)
}

//ChackHexadecimalZeros :
func ChackHexadecimalZeros(hexadecimal []byte, ze string) (bool, string) {
	encodeString := hex.EncodeToString(hexadecimal)
	if !strings.HasPrefix(encodeString, ze) {
		return false, ""
	}
	result := ze
	triesZeros := ze
	for {
		if strings.HasPrefix(encodeString, triesZeros) {
			result = triesZeros
		} else {
			return true, result
		}
		triesZeros += "0"
	}
}

//BestZero :
type BestZero struct {
	Text      string
	Nonce     int
	TextNonce string
	Zeros     string
	Checksum  string
}
