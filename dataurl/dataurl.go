package dataurl

import (
	"encoding/base64"
	"strings"
)

func Decode(str string) (data []byte, err error) {
	return base64.StdEncoding.DecodeString(strings.SplitN(str, ",", 2)[1])
}

// func Encode() {

// }
