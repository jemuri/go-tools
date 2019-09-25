package encr

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 encryption
func MD5(source string) string {
	h:= md5.New()

	h.Write([]byte(source))

	return hex.EncodeToString(h.Sum(nil))
}