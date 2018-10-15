package encryptfilessymmetric_test

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"testing"

	"github.com/tarekbadrshalaan/goStuff/encryptfilessymmetric"
)

//!+test
//go test -v
func TestEncryptDecrypt(t *testing.T) {
	var tests = []struct {
		key string
		raw string
	}{
		{"", ""},
		{"key1", "test"},
		{"test", "Clear is better than clever"},
		{"Lorem Ipsum", "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."},
	}
	for i, test := range tests {
		writer := new(bytes.Buffer)
		reader := bytes.NewBufferString(test.raw)

		// Encrypt
		err := encryptfilessymmetric.Encrypt(reader, writer, test.key)
		if err != nil {
			t.Errorf("TestEncrypt Encrypt error %v", err)
		}
		// Decrypt
		result := new(bytes.Buffer)
		err = encryptfilessymmetric.Decrypt(writer, result, test.key)
		if err != nil {
			t.Errorf("TestEncrypt Decrypt error %v", err)
		}

		resultstr := result.String()
		if resultstr != test.raw {
			t.Errorf("TestEncrypt index : %d ;result not as expected\nexpected:%v\nactual:%v", i, test.raw, resultstr)
		}
	}
}

func TestRandom(t *testing.T) {

	rawbuff := make([]byte, 50000)
	rand.Read(rawbuff)

	keybuff := make([]byte, 50)
	rand.Read(keybuff)
	key := base64.StdEncoding.EncodeToString(keybuff)

	writer := new(bytes.Buffer)
	reader := bytes.NewBuffer(rawbuff)

	// Encrypt
	err := encryptfilessymmetric.Encrypt(reader, writer, key)
	if err != nil {
		t.Errorf("TestRandom Encrypt error %v", err)
	}
	// Decrypt
	result := new(bytes.Buffer)
	err = encryptfilessymmetric.Decrypt(writer, result, key)
	if err != nil {
		t.Errorf("TestRandom Decrypt error %v", err)
	}

	if result.String() != string(rawbuff) {
		t.Errorf("TestRandom result not as expected\nexpected:%s\nactual:%s", result, rawbuff)
	}
}

//!-tests

//!+bench
//go test -v  -bench=.
func BenchmarkEncryptDecrypt(b *testing.B) {
	key := "test"
	raw := "Clear is better than clever"
	for index := 0; index < b.N; index++ {
		writer := new(bytes.Buffer)
		reader := bytes.NewBufferString(raw)
		err := encryptfilessymmetric.Encrypt(reader, writer, key)
		if err != nil {
			b.Errorf("BenchmarkEncryptDecrypt Encrypt error %v", err)
		}
		err = encryptfilessymmetric.Decrypt(writer, reader, key)
		if err != nil {
			b.Errorf("BenchmarkEncryptDecrypt Decrypt error %v", err)
		}
		result := reader.String()
		if result != raw {
			b.Errorf("BenchmarkEncryptDecrypt index : %d ;result not as expected\nexpected:%v\nactual:%v", index, raw, result)
		}
	}
}

//!-bench
