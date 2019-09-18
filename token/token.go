package token

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func Generate(source string) string {
	curTime := time.Now().Unix()
	h:= md5.New()

	h.Write([]byte(fmt.Sprintf("%s%d",source,curTime)))

	return hex.EncodeToString(h.Sum(nil))
}
