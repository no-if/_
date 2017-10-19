package gbk

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

func Encode(src []byte) ([]byte, error) {
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader(src), simplifiedchinese.GBK.NewEncoder()))
}
func Decode(src []byte) ([]byte, error) {
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader(src), simplifiedchinese.GBK.NewDecoder()))
}
