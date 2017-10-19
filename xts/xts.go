package xts

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"golang.org/x/crypto/xts"
	"math"
)

const (
	BlockSize = 16
)

func Cipher(key string) (*xts.Cipher, error) {
	var data, err = hex.DecodeString(key)
	if err == nil {
		return xts.NewCipher(aes.NewCipher, data)
	}
	return new(xts.Cipher), err
}

func Encode(cipher *xts.Cipher, src []byte, sector uint64) (dst []byte) {
	var data []byte = make([]byte, int(math.Ceil(float64(len(src))/16)*16))
	copy(data, src)
	dst = make([]byte, len(data))
	cipher.Encrypt(dst, data, sector)
	return
}

func Decode(cipher *xts.Cipher, src []byte, sector uint64) (dst []byte) {
	var l = len(src)
	if l%16 != 0 {
		return
	}
	dst = make([]byte, l)
	cipher.Decrypt(dst, src, sector)
	dst = bytes.TrimRightFunc(dst, func(r rune) bool {
		return r == 0
	})
	return
}
