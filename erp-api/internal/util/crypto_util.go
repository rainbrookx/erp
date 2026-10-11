package util

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Crypto(s string) string {
	data := []byte(s)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
