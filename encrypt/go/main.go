package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"

	pkcs7pad "main/pkcs7pad"
)

func main() {

	msgByte := []byte(os.Args[1])
	if len(msgByte)%16 != 0 {
		msgByte = pkcs7pad.Pad(msgByte, 16)
	}

	msg := base64.StdEncoding.EncodeToString(msgByte)

	key := base64.StdEncoding.EncodeToString([]byte(os.Getenv("Key")))
	iv := base64.StdEncoding.EncodeToString([]byte(os.Getenv("IV")))

	result := encryptAction(msg, key, iv)
	fmt.Println(result)
}

func encryptAction(msg string, key string, iv string) string {

	msgAsBytes, _ := base64.StdEncoding.DecodeString(msg)
	ivAsByte, _ := base64.StdEncoding.DecodeString(iv)
	keyAsByte, _ := base64.StdEncoding.DecodeString(key)

	block, _ := aes.NewCipher(keyAsByte)
	cipherText := make([]byte, len(msgAsBytes))

	stream := cipher.NewCBCEncrypter(block, ivAsByte)
	stream.CryptBlocks(cipherText, msgAsBytes)

	return base64.StdEncoding.EncodeToString(cipherText)
}
