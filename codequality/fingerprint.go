package codequality

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func Fingerprint(element Element) string {
	j, _ := json.Marshal(element)
	sum := sha256.Sum256(j)
	return hex.EncodeToString(sum[:])
}
