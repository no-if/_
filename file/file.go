package file

import (
	"os"
)

func Exist(file_name string) bool {
	_, err := os.Lstat(file_name)
	return err == nil || os.IsExist(err)
}
