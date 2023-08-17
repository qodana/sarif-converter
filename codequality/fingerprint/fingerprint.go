package fingerprint

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"sarif-converter/codequality/element"
)

func Fingerprint(element element.Element) string {
	j, _ := json.Marshal(element)
	sum := sha256.Sum256(j)
	return hex.EncodeToString(sum[:])
}
