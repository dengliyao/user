package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5string(text string) string {
	return fmt.Sprintf("%X", md5.Sum([]byte(text)))
}
