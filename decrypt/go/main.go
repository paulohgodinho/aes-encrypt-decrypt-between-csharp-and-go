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
	msg := os.Args[1]
	key := base64.StdEncoding.EncodeToString([]byte(os.Getenv("Key")))
	iv := base64.StdEncoding.EncodeToString([]byte(os.Getenv("IV")))

	result := decryptAction(msg, key, iv)
	fmt.Println(result)
}

func decryptAction(msg string, key string, iv string) string {

	msgAsBytes, _ := base64.StdEncoding.DecodeString(msg)
	ivAsByte, _ := base64.StdEncoding.DecodeString(iv)
	keyAsByte, _ := base64.StdEncoding.DecodeString(key)

	block, _ := aes.NewCipher(keyAsByte)
	cipherText := make([]byte, len(msgAsBytes))

	stream := cipher.NewCBCDecrypter(block, ivAsByte)
	stream.CryptBlocks(cipherText, msgAsBytes)

	removedPadding, _ := pkcs7pad.Unpad(cipherText)
	return string(removedPadding)
}
