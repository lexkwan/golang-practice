package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"flag"
	"fmt"
)

// DesEncryption encryption text with DES method
func DesEncryption(key, iv, plainText []byte) ([]byte, error) {

	block, err := des.NewCipher(key)

	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData := PKCS5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return cryted, nil
}

// DesDecryption decrypte the text
func DesDecryption(key, iv, cipherText []byte) ([]byte, error) {

	block, err := des.NewCipher(key)

	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

// PKCS5Padding pkcs5 padding
func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

//PKCS5UnPadding unpadding text
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

var originalText, key, iv *string

func init() {
	originalText = flag.String("string", "", "The original text you want to encrypte")
	key = flag.String("key", "abcdef", "des encryption private key")
	iv = flag.String("iv", "123456", "des encryption iv")
}

func main() {

	flag.Parse()

	mytext := []byte(*originalText)

	k := []byte(*key)
	i := []byte(*iv)

	cryptoText, _ := DesEncryption(k, i, mytext)
	fmt.Println(base64.StdEncoding.EncodeToString(cryptoText))

	// decryptedText, _ := DesDecryption(key, iv, cryptoText)
	// fmt.Println(string(decryptedText))

}
