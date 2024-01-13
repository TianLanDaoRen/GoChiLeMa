package encrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

type Encrypto struct{}

func NewEncrypto() *Encrypto {
	return &Encrypto{}
}

func (e *Encrypto) Encrypt(plaintext string) string {
	// ASE-256-CBC
	bKey := []byte("g4t8g2b1hj4yu5hrt8d4g25iuj18g59k")
	bIv := []byte("s83jf8d4g2h1j8d4")
	bPlaintext := e.PKCS5Padding([]byte(plaintext), aes.BlockSize)
	block, err := aes.NewCipher(bKey)
	if err != nil {
		return ""
	}
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIv)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext)
}

func (e *Encrypto) PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (e *Encrypto) Decrypt(ciphertext string) string {
	// ASE-256-CBC
	bKey := []byte("g4t8g2b1hj4yu5hrt8d4g25iuj18g59k")
	bIv := []byte("s83jf8d4g2h1j8d4")
	// base64 decode
	cipherTextDecoded, err := hex.DecodeString(ciphertext)
	if err != nil {
		return ""
	}
	block, err := aes.NewCipher(bKey)
	if err != nil {
		return ""
	}
	mode := cipher.NewCBCDecrypter(block, bIv)
	mode.CryptBlocks([]byte(cipherTextDecoded), []byte(cipherTextDecoded))
	return string(cipherTextDecoded)
}
