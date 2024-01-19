package encrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

//go:embed key.json
var key string

type Encrypto struct {
	Key string `json:"key"`
	Iv  string `json:"Iv"`
}

func NewEncrypto() *Encrypto {
	// Unmarshal key string to json
	var encrypto Encrypto
	err := json.Unmarshal([]byte(key), &encrypto)
	if err != nil {
		fmt.Println(err)
		return &Encrypto{}
	}
	return &encrypto
}

func (e *Encrypto) Encrypt(plaintext string) string {
	// ASE-256-CBC
	bKey := []byte(e.Key)
	bIv := []byte(e.Iv)
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
	bKey := []byte(e.Key)
	bIv := []byte(e.Iv)
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
