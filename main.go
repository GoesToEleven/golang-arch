package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	msg := "This is totally fun get hands-on and learning it from the ground up. Thank you for sharing this info with me and helping me learn!"
	encoded := encode(msg)
	fmt.Println("ENCODED MSG", encoded)

	s, err := decode(encoded)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("DECODED MSG", s)
}

func encode(msg string) string {
	return base64.URLEncoding.EncodeToString([]byte(msg))
}

func decode(encoded string) (string, error) {
	s, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return "", fmt.Errorf("couldn't decode string %w", err)
	}
	return string(s), nil
}
