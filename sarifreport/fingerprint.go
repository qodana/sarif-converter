package sarifreport

import (
	"codequality-converter/codequality"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type Element struct {
	location    codequality.CodeQualityLocation
	description *string
}

func Fingerprint(location codequality.CodeQualityLocation, description *string) string {
	output, _ := json.Marshal(Element{location: location, description: description})
	h := sha256.New()
	return hex.EncodeToString(h.Sum(output))
}
