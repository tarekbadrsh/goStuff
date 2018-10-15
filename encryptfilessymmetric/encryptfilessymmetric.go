// encryptfilessymmetric : package use to Encrypt/Decrypt large Data
// Basically it was created to be used in large files
// but it will work with any kind of  io.Reader&w io.Writer

package encryptfilessymmetric

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"io"
	"os"

	"github.com/pkg/errors"
)

// NewStream : return new stream or error.
func NewStream(keytext string) (cipher.Stream, error) {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	hashkey := sha256.Sum256([]byte(keytext))
	key := hashkey[:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrap(err, "could not create NewCipher")
	}

	// If the key is unique for each ciphertext, then it's ok to use a zero
	// IV.
	var iv [aes.BlockSize]byte
	return cipher.NewOFB(block, iv[:]), nil
}

// Encrypt : encrypt reader and copy to writer with password.
func Encrypt(r io.Reader, w io.Writer, keytext string) error {
	stream, err := NewStream(keytext)
	if err != nil {
		return errors.Wrap(err, "could not create NewStream")
	}
	writer := &cipher.StreamWriter{S: stream, W: w}
	// Copy the input file to the output file, encrypting as we go.
	if _, err := io.Copy(writer, r); err != nil {
		return errors.Wrap(err, "could not Copy writer, inFile")
	}
	return nil
}

// Decrypt : decrypt reader and copy to writer with password.
func Decrypt(r io.Reader, w io.Writer, keytext string) error {
	stream, err := NewStream(keytext)
	if err != nil {
		return errors.Wrap(err, "could not create NewStream")
	}

	reader := &cipher.StreamReader{S: stream, R: r}
	// Copy the input file to the output file, decrypting as we go.
	if _, err := io.Copy(w, reader); err != nil {
		return errors.Wrap(err, "could not Copy")
	}
	return nil
}

// EncryptFile in/out file with password.
func EncryptFile(inFileName, outFileName, keytext string) error {
	inFile, err := os.Open(inFileName)
	if err != nil {
		return errors.Wrapf(err, "could not Open inFileName:%v", inFileName)
	}
	defer inFile.Close()

	outFile, err := os.OpenFile(outFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return errors.Wrapf(err, "could not Open outFileName:%v", outFileName)
	}
	defer outFile.Close()

	err = Encrypt(inFile, outFile, keytext)
	if err != nil {
		return errors.Wrap(err, "could not Encrypt")
	}
	return nil
}

// DecryptFile in/out file with password.
func DecryptFile(inFileName, outFileName, keytext string) error {
	inFile, err := os.Open(inFileName)
	if err != nil {
		return errors.Wrapf(err, "could not Open inFileName:%v", inFileName)
	}
	defer inFile.Close()

	outFile, err := os.OpenFile(outFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return errors.Wrapf(err, "could not Open outFileName:%v", outFileName)
	}
	defer outFile.Close()

	err = Decrypt(inFile, outFile, keytext)
	if err != nil {
		return errors.Wrap(err, "could not Decrypt")
	}
	return nil
}
