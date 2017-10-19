package hmac

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"hash"
)

func encode(h func() hash.Hash, src []byte, key string) (dst string) {
	mac := hmac.New(h, []byte(key))
	mac.Write(src)
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func Sha1(src []byte, key string) (dst string) {
	return encode(sha1.New, src, key)
}

func Md5(src []byte, key string) (dst string) {
	return encode(md5.New, src, key)
}
