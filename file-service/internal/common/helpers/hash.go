package helpers

import (
	"crypto/sha256"
	"fmt"
)

func GenerateHash(data []byte) string {
	hash := sha256.Sum256(data)

	return fmt.Sprintf("%x", hash)
}
