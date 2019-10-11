package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func main() {
	key := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	encrypted, err := encrypt([]byte("Hello World!"), key)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(encrypted))

	decrypted, err := encrypt(encrypted, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decrypted))
}

func encrypt(message, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("Error in encryption: %w", err)
	}

	stream := cipher.NewOFB(block, make([]byte, aes.BlockSize))
	buf := &bytes.Buffer{}
	wtr := cipher.StreamWriter{
		S: stream,
		W: buf,
	}

	_, err = wtr.Write(message)
	if err != nil {
		return nil, fmt.Errorf("Error in encryption: %w", err)
	}

	return buf.Bytes(), err
}
