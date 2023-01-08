package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(text string) string {
	hash := sha256.New()
	hash.Write([]byte(text))

	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}
