package hashit

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func Hasher(password string) {
	h := sha1.New()
	h.Write([]byte(password))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	fmt.Println(sha1_hash)
}
